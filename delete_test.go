// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian_test

import (
	"testing"

	"cunicu.li/go-feitian-oath"
)

func TestDelete(t *testing.T) {
	withCard(t, false, func(_ *testing.T, _ *feitian.Card) {
		// require := require.New(t)
	})
}
