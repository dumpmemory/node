/*
 * Copyright (C) 2026 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package pingpong

import (
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/node/identity"
	"github.com/mysteriumnetwork/node/identity/registry"
	"github.com/mysteriumnetwork/payments/observer"
	"github.com/stretchr/testify/assert"
)

// TestPromiseSettler_resyncActiveChannel_FetchesUnlockedIdentity encodes the fix for ClickUp
// 869dwcqn9 (UI "Unsettled" ghost balance under settlement rate limits).
//
// Root cause (verified): channelProvider.Fetch() — the only thing that refreshes the local
// Channel.Settled value used by UnsettledBalance() — runs only inside the settlement flow
// (post-settlement resync at hermes_promise_settler.go:1145). When Hermes rate-limits
// settlement, settle() returns early (UpdatePromiseFee returns HermesErrorResponse(ErrTooManyRequests)
// with no retry), so the resync Fetch() never runs. Meanwhile the Transactor's own cron settles
// on-chain, advancing Settled, but the node never re-reads it, so UnsettledBalance grows.
//
// Fix: a periodic background sync, decoupled from settlement, that fetches the true channel
// balance for the single active (unlocked) identity in effect — multi-identity is legacy. This
// test exercises the seam resyncActiveChannel(chainID) directly: no settlement, no promise event,
// no timing/sleep.
func TestPromiseSettler_resyncActiveChannel_FetchesUnlockedIdentity(t *testing.T) {
	channelProvider := &countingChannelProvider{}
	activeHermes := common.HexToAddress("0x000000000000000000000000000000000000beef")
	obs := &activeHermesObserver{hermesByChain: observer.HermesesResponse{
		0: []observer.HermesResponse{{HermesAddress: activeHermes}},
	}}

	settler := newResyncTestSettler(channelProvider, obs, &mockCurrentIdentity{id: mockID, unlocked: true})

	// The active identity is registered — the node must keep its balance in sync independent of
	// any settlement happening.
	settler.currentState[mockID] = settlementState{
		registered:       true,
		settleInProgress: map[common.Address]struct{}{},
	}

	settler.resyncActiveChannel(settler.chainID())

	assert.Equal(t, 1, channelProvider.fetchCount(), "expected one Fetch for the single active identity x one active hermes")
	assert.Contains(t, channelProvider.fetchedIdentities(), mockID, "active identity must be resynced regardless of settlement")
}

// TestPromiseSettler_resyncActiveChannel_NoUnlockedIdentity verifies that when there is no active
// identity in effect, the resync does nothing — no infra calls are made.
func TestPromiseSettler_resyncActiveChannel_NoUnlockedIdentity(t *testing.T) {
	channelProvider := &countingChannelProvider{}
	obs := &activeHermesObserver{hermesByChain: observer.HermesesResponse{
		0: []observer.HermesResponse{{HermesAddress: common.HexToAddress("0x1")}},
	}}

	settler := newResyncTestSettler(channelProvider, obs, &mockCurrentIdentity{unlocked: false})
	settler.currentState[mockID] = settlementState{registered: true, settleInProgress: map[common.Address]struct{}{}}

	settler.resyncActiveChannel(settler.chainID())

	assert.Equal(t, 0, channelProvider.fetchCount(), "no active identity -> no resync, no infra calls")
}

// TestPromiseSettler_clampResyncInterval enforces the hard infra floor: the resync interval can
// never drop below defaultChannelResyncInterval (15m), regardless of configuration, protecting
// shared Hermes/blockchain infra from a large node fleet.
func TestPromiseSettler_clampResyncInterval(t *testing.T) {
	assert.Equal(t, defaultChannelResyncInterval, clampResyncInterval(0), "unset -> floor")
	assert.Equal(t, defaultChannelResyncInterval, clampResyncInterval(time.Minute), "below floor -> floor")
	assert.Equal(t, defaultChannelResyncInterval, clampResyncInterval(defaultChannelResyncInterval), "exactly floor -> floor")
	assert.Equal(t, 20*time.Minute, clampResyncInterval(20*time.Minute), "above floor -> unchanged")
	assert.GreaterOrEqual(t, int64(clampResyncInterval(time.Second)), int64(15*time.Minute), "floor is at least 15m")
}

func newResyncTestSettler(cp hermesChannelProvider, obs observerApi, ci currentIdentityProvider) *hermesPromiseSettler {
	mrsp := &mockRegistrationStatusProvider{
		identities: map[string]mockRegistrationStatus{
			mockChainIdentity: {status: registry.Registered},
		},
	}
	fac := &mockHermesCallerFactory{}
	return NewHermesPromiseSettler(
		&mockTransactor{},
		&mockHermesPromiseStorage{},
		&mockPayAndSettler{},
		&mockAddressProvider{},
		fac.Get,
		&mockHermesURLGetter{},
		cp,
		&mockProviderChannelStatusProvider{},
		mrsp,
		identity.NewMockKeystore(),
		ci,
		&settlementHistoryStorageMock{},
		&mockPublisher{},
		obs,
		newMockAddressStorage(),
		cfg,
	)
}

// mockCurrentIdentity is the single active (unlocked) identity in effect for the runtime.
type mockCurrentIdentity struct {
	id       identity.Identity
	unlocked bool
}

func (m *mockCurrentIdentity) GetUnlockedIdentity() (identity.Identity, bool) {
	return m.id, m.unlocked
}

// countingChannelProvider is a hermesChannelProvider that records Fetch calls so we can assert
// the node resyncs channel balances out-of-band from the settlement flow.
type countingChannelProvider struct {
	mu      sync.Mutex
	count   int
	fetched []identity.Identity
}

func (c *countingChannelProvider) Get(chainID int64, _ identity.Identity, _ common.Address) (HermesChannel, bool) {
	return HermesChannel{}, true
}

func (c *countingChannelProvider) Fetch(_ int64, id identity.Identity, _ common.Address) (HermesChannel, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	c.fetched = append(c.fetched, id)
	return HermesChannel{}, nil
}

func (c *countingChannelProvider) fetchCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func (c *countingChannelProvider) fetchedIdentities() []identity.Identity {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]identity.Identity, len(c.fetched))
	copy(out, c.fetched)
	return out
}

// activeHermesObserver returns a configurable set of active hermeses per chain.
type activeHermesObserver struct {
	hermesByChain observer.HermesesResponse
}

func (o *activeHermesObserver) GetHermeses(_ *observer.HermesFilter) (observer.HermesesResponse, error) {
	return o.hermesByChain, nil
}

func (o *activeHermesObserver) GetHermesData(_ int64, _ common.Address) (*observer.HermesResponse, error) {
	return nil, nil
}
