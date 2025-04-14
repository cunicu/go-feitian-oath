// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian

import (
	"encoding/binary"

	iso "cunicu.li/go-iso7816"
	"cunicu.li/go-iso7816/encoding/tlv"
)

// Put programs a OTP credential.
func (c *Card) Put(slot Slot, name string, secret []byte, alg Algorithm, kind Kind, digits int, counter uint32) error {
	if err := checkName(name); err != nil {
		return err
	} else if err := checkSlot(slot); err != nil {
		return err
	} else if err := checkDigits(digits); err != nil {
		return err
	}

	data, err := tlv.EncodeSimple(
		tlv.New(tagKey, byte(alg), byte(kind), byte(digits), secret),
		tlv.New(tagName, name),
		tlv.New(tagTouch, byte(tagTouch)),
		tlv.New(tagIMF, binary.BigEndian.AppendUint32(nil, counter)))
	if err != nil {
		return err
	}

	// Really weird quirk?
	data[1]--

	_, err = c.Send(&iso.CAPDU{
		Ins:  insPut,
		P1:   0x00,
		P2:   byte(slot),
		Data: data,
	})

	return err
}
