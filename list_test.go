// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cunicu.li/go-feitian-oath"
)

func TestList(t *testing.T) {
	withCard(t, true, func(t *testing.T, c *feitian.Card) {
		require := require.New(t)

		items, err := c.List()
		require.NoError(err)
		require.Len(items, 0)

		err = c.Put(feitian.Slot1, "test", testSecretSHA1, feitian.SHA1, feitian.TOTP, 6, 0)
		require.NoError(err)

		items, err = c.List()
		require.NoError(err)
		require.Len(items, 1)
	})
}
