// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cunicu.li/go-feitian-oath"
)

func TestPut(t *testing.T) {
	withCard(t, false, func(t *testing.T, c *feitian.Card) {
		require := require.New(t)

		for _, vs := range vectors {
			for _, v := range vs {
				err := c.Put(feitian.Slot1, v.Name, v.Secret, v.Algorithm, v.Kind, v.Digits, v.Counter)
				require.NoError(err)

				items, err := c.List()
				require.NoError(err)
				require.Len(items, 1)

				require.Equal(items[0].Algorithm, v.Algorithm)
				require.Equal(items[0].Kind, v.Kind)
				require.Equal(items[0].Name, v.Name)
				require.Equal(items[0].Slot, feitian.Slot1)
			}
		}
	})
}
