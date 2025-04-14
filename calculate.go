// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian

import (
	"encoding/binary"
	"errors"
	"time"

	iso "cunicu.li/go-iso7816"
	"cunicu.li/go-iso7816/encoding/tlv"
)

var ErrMissingResponse = errors.New("missing response")

func (c *Card) Calculate(slot Slot, name string) (Code, error) {
	challenge := ChallengeTOTP(c.Clock(), c.Timestep)
	return c.CalculateWithChallenge(slot, name, challenge, false)
}

// CalculateWithChallenge the OTP value.
func (c *Card) CalculateWithChallenge(slot Slot, name string, challenge []byte, truncate bool) (Code, error) {
	if err := checkSlot(slot); err != nil {
		return Code{}, err
	} else if err := checkName(name); err != nil {
		return Code{}, err
	}

	data, err := tlv.EncodeSimple(
		tlv.New(tagChallenge, challenge),
		tlv.New(tagName, name),
	)
	if err != nil {
		return Code{}, err
	}

	var p1 byte = 0x00
	if truncate {
		p1 = 0x01
	}

	resp, err := c.Send(&iso.CAPDU{
		Ins:  insCalculate,
		P1:   p1,
		P2:   byte(slot),
		Data: data,
	})
	if err != nil {
		return Code{}, err
	}

	tvs, err := tlv.DecodeSimple(resp)
	if err != nil {
		return Code{}, err
	}

	for _, tv := range tvs {
		switch tv.Tag {
		case tagResponse, tagTResponse:
			return Code{
				Digits:    int(tv.Value[0]),
				Digest:    tv.Value[1:],
				Truncated: tv.Tag == tagTResponse,
			}, nil
		}
	}

	return Code{}, ErrMissingResponse
}

func ChallengeTOTP(t time.Time, ts time.Duration) []byte {
	counter := t.Unix() / int64(ts.Seconds())
	return binary.BigEndian.AppendUint64(nil, uint64(counter)) //nolint:gosec
}
