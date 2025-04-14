// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cunicu.li/go-feitian-oath"
)

func TestSwap(t *testing.T) {
	withCard(t, false, func(t *testing.T, c *feitian.Card) {
		require := require.New(t)

		err := c.Swap()
		require.NoError(err)
	})
}
