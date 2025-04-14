// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian

import (
	"errors"
	"fmt"

	iso "cunicu.li/go-iso7816"
	"cunicu.li/go-iso7816/encoding/tlv"
)

var ErrSlotNotConfigured = errors.New("slot has not credential configured")

// Set active credential (Deprecated).
func (c *Card) SetDefault(slot Slot, name string) error {
	if err := checkName(name); err != nil {
		return err
	} else if err := checkSlot(slot); err != nil {
		return err
	}

	data, err := tlv.EncodeSimple(tlv.New(tagName, name))
	if err != nil {
		return fmt.Errorf("failed to encode slot name: %w", err)
	}

	_, err = c.Send(&iso.CAPDU{
		Ins:  insSetDefault,
		P1:   0x00,
		P2:   byte(slot),
		Data: data,
	})

	return err
}

// Get active credential (Deprecated).
func (c *Card) Default(slot Slot) (string, error) {
	if err := checkSlot(slot); err != nil {
		return "", err
	}

	resp, err := c.Send(&iso.CAPDU{
		Ins: insGetDefault,
		P1:  0x00,
		P2:  byte(slot),
	})
	if err != nil {
		return "", err
	}

	if len(resp) == 0 {
		return "", ErrSlotNotConfigured
	} else if len(resp) < 3 {
		return "", iso.ErrWrongLength
	}

	return string(resp[3:]), nil
}
