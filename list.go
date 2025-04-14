// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian

import (
	"cunicu.li/go-iso7816"
	"cunicu.li/go-iso7816/encoding/tlv"
)

type ListItem struct {
	Name      string
	Slot      Slot
	Algorithm Algorithm
	Kind      Kind
}

// List gets OTP credentials of all slots.
func (c *Card) List() ([]ListItem, error) {
	items := []ListItem{}

	for _, slot := range []Slot{Slot1, Slot2, SlotDefault} {
		resp, err := c.Send(&iso7816.CAPDU{
			Ins: insList,
			P1:  0x00,
			P2:  byte(slot),
		})
		if err != nil {
			return nil, err
		}

		tvs, err := tlv.DecodeSimple(resp)
		if err != nil {
			return nil, err
		}

		for _, tv := range tvs {
			v := tv.Value

			item := ListItem{
				Kind:      Kind(v[0] & 0xF0),
				Algorithm: Algorithm(v[0] & 0x0F),
				Name:      string(v[1:]),
			}

			items = append(items, item)
		}
	}

	return items, nil
}
