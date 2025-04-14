// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cunicu.li/go-feitian-oath"
)

func TestLanguage(t *testing.T) {
	withCard(t, true, func(t *testing.T, c *feitian.Card) {
		require := require.New(t)

		lang, err := c.Language()
		require.NoError(err)
		require.Equal(feitian.LangFrench, lang)

		err = c.SetLanguage(feitian.LangEnglish)
		require.NoError(err)

		lang, err = c.Language()
		require.NoError(err)
		require.Equal(feitian.LangEnglish, lang)
	})
}
