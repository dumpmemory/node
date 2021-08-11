/*
 * Copyright (C) 2021 The "MysteriumNetwork/node" Authors.
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

package nat

import (
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/mysteriumnetwork/node/config"
)

type PortProvider interface {
	PreparePorts() (ports []int, release func(), start StartPorts, err error)
}

var traversalOptions map[string]func() PortProvider = map[string]func() PortProvider{
	"manual":       NewManualPortProvider,
	"upnp":         NewUPnPPortProvider,
	"holepunching": NewNatHolePunchingPortProvider,
}

func OrderedPortProviders() (list []PortProvider) {
	methods := strings.Split(config.GetString(config.FlagTraversal), ",")

	for _, m := range methods {
		if t, ok := traversalOptions[m]; ok {
			list = append(list, t())
		} else {
			log.Warn().Msgf("Unsupported traversal method %s, ignoring it", m)
		}
	}

	return list
}
