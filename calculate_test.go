// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cunicu.li/go-feitian-oath"
)

func TestCalculateTOTP(t *testing.T) {
	withCardAndVectors(t, vectorsTOTP, true, func(t *testing.T, c *feitian.Card, v vector) {
		require := require.New(t)

		err := c.Put(feitian.Slot1, v.Name, v.Secret, v.Algorithm, v.Kind, v.Digits, v.Counter)
		require.NoError(err)

		code, err := c.CalculateWithChallenge(feitian.Slot1, v.Name, feitian.ChallengeTOTP(v.Time, feitian.DefaultTimeStep), false)
		require.NoError(err)
		require.Equal(v.Digits, code.Digits)
		require.Equal(v.Code, code.OTP())
	})
}

func TestCalculateHOTPCounterIncrement(t *testing.T) {
	withCard(t, true, func(t *testing.T, c *feitian.Card) {
		require := require.New(t)

		v := vectorsHOTP[0]

		err := c.Put(feitian.Slot1, v.Name, v.Secret, v.Algorithm, v.Kind, v.Digits, v.Counter)
		require.NoError(err)

		for _, ev := range vectorsHOTP[:10] {
			code, err := c.Calculate(feitian.Slot1, v.Name)
			require.NoError(err)
			require.Equal(ev.Code, code.OTP())
			require.Equal(v.Digits, code.Digits)
			require.False(code.Truncated)
		}
	})
}

func TestCalculateHOTPCounterInit(t *testing.T) {
	withCardAndVectors(t, vectorsHOTP, true, func(t *testing.T, c *feitian.Card, v vector) {
		require := require.New(t)

		err := c.Put(feitian.Slot1, v.Name, v.Secret, v.Algorithm, v.Kind, v.Digits, 0)
		require.NoError(err)

		code, err := c.Calculate(feitian.Slot1, v.Name)
		require.NoError(err)
		require.Equal(v.Code, code.OTP())
		require.Equal(v.Digits, code.Digits)
		require.False(code.Truncated)
	})
}

func TestCalculateWithChallenge(t *testing.T) {
	withCardAndVectors(t, vectorsChalResp, true, func(t *testing.T, c *feitian.Card, v vector) {
		require := require.New(t)

		err := c.Put(feitian.Slot1, v.Name, v.Secret, v.Algorithm, v.Kind, v.Digits, 0)
		require.NoError(err)

		code, err := c.CalculateWithChallenge(feitian.Slot1, v.Name, v.Challenge, false)
		require.NoError(err)
		require.Equal(v.Digits, code.Digits)
		require.Equal(v.Hash, code.Digest)
	})
}

func TestCalculateStaticPassword(t *testing.T) {
	withCard(t, true, func(t *testing.T, c *feitian.Card) {
		require := require.New(t)

		pass := []byte("my static password")

		err := c.Put(feitian.Slot1, "test", pass, feitian.SHA1, feitian.StaticPassword, 6, 0)
		require.NoError(err)

		code, err := c.CalculateWithChallenge(feitian.Slot1, "test", nil, false)
		require.NoError(err)
		require.Equal(6, code.Digits)
		require.Equal(pass, code.Digest)
	})
}
