// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cunicu.li/go-feitian-oath"
)

func TestDefault(t *testing.T) {
	withCard(t, true, func(t *testing.T, c *feitian.Card) {
		require := require.New(t)

		_, err := c.Default(feitian.Slot1)
		require.ErrorIs(err, feitian.ErrSlotNotConfigured)
	})
}
