/*
 * Copyright (C) 2025 The "MysteriumNetwork/node" Authors.
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

package streams

import (
	"github.com/quic-go/quic-go"
	"github.com/rs/zerolog/log"
)

// QuicConnection represents QUIC connection.
type QuicConnection struct {
	quic.Connection
	Listener *quic.Listener
}

// Close closes connection and listener.
func (c QuicConnection) Close() error {
	if c.Listener != nil {
		if err := c.Listener.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close QUIC listener")
		}
	}

	return c.CloseWithError(0, "")
}
