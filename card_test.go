// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian_test

import (
	"testing"

	iso "cunicu.li/go-iso7816"
	"cunicu.li/go-iso7816/filter"
	"cunicu.li/go-iso7816/test"
	"github.com/stretchr/testify/require"

	"cunicu.li/go-feitian-oath"
)

func withCard(t *testing.T, reset bool, cb func(t *testing.T, c *feitian.Card)) {
	test.WithCard(t, filter.IsFeitian, func(t *testing.T, c *iso.Card) {
		require := require.New(t)

		fc, err := feitian.NewCard(c)
		require.NoError(err)

		err = fc.Select()
		require.NoError(err)

		if reset {
			err = fc.Reset()
			require.NoError(err)
		}

		cb(t, fc)
	})
}

func withCardAndVectors(t *testing.T, vectors []vector, reset bool, cb func(t *testing.T, c *feitian.Card, v vector)) {
	withCard(t, reset, func(t *testing.T, c *feitian.Card) {
		for _, vector := range vectors {
			cb(t, c, vector)
		}
	})
}
