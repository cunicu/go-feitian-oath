// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian

import (
	"errors"
	"fmt"
	"time"

	iso "cunicu.li/go-iso7816"
	"cunicu.li/go-iso7816/devices/feitian"
	"cunicu.li/go-iso7816/encoding/tlv"
)

type Slot byte

const (
	Slot1       Slot = 0x00
	Slot2       Slot = 0x01
	SlotDefault Slot = 0xF0
)

const DefaultTimeStep = 30 * time.Second

type Algorithm byte

const (
	SHA1   Algorithm = 0x01
	SHA256 Algorithm = 0x02
)

const (
	tagName       tlv.Tag = 0x51
	tagNameList   tlv.Tag = 0x52
	tagKey        tlv.Tag = 0x53
	tagChallenge  tlv.Tag = 0x54
	tagResponse   tlv.Tag = 0x55
	tagProperty   tlv.Tag = 0x58
	tagVersion    tlv.Tag = 0x59
	tagIMF        tlv.Tag = 0x5A // Initial Moving Factor (counter value for HOTP)
	tagAlgorithm  tlv.Tag = 0x5B
	tagTouch      tlv.Tag = 0x5C
	tagTResponse  tlv.Tag = 0x76 // Truncated
	tagNoResponse tlv.Tag = 0x77
)

const (
	insSetCode         iso.Instruction = 0x02
	insReset           iso.Instruction = 0x07
	insDelete          iso.Instruction = 0x08
	insPut             iso.Instruction = 0x09
	insList            iso.Instruction = 0x17
	insValidate        iso.Instruction = 0x19
	insCalculateAll    iso.Instruction = 0x1A
	insSendRemaining   iso.Instruction = 0x1B
	insCalculate       iso.Instruction = 0xA2
	insLanguage        iso.Instruction = 0xA7
	insSendRemainingFT iso.Instruction = 0xC0
	insApplication     iso.Instruction = 0xE1
	insSetDefault      iso.Instruction = 0xE5
	insGetDefault      iso.Instruction = 0xE6
	insSwapSlot        iso.Instruction = 0xE7
)

type AppState byte

const (
	ON  AppState = 0x01
	OFF AppState = 0x00
)

type Kind byte

const (
	HOTP              Kind = 0x10
	TOTP              Kind = 0x20
	StaticPassword    Kind = 0x30
	ChallengeResponse Kind = 0x40
)

var (
	ErrNameTooShort  = errors.New("name is too short")
	ErrNameTooLong   = errors.New("name is too long")
	ErrInvalidSlot   = errors.New("invalid slot")
	ErrInvalidDigits = errors.New("number of digits must be either 6 or 8")
)

type Card struct {
	*feitian.Card

	Clock    func() time.Time
	Timestep time.Duration

	tx *iso.Transaction
}

// NewCard initializes a new card.
func NewCard(pcscCard iso.PCSCCard) (*Card, error) {
	isoCard := iso.NewCard(pcscCard)
	isoCard.InsGetRemaining = insSendRemaining

	tx, err := isoCard.NewTransaction()
	if err != nil {
		return nil, fmt.Errorf("failed to initiate transaction: %w", err)
	}

	return &Card{
		Card:     &feitian.Card{Card: isoCard},
		Clock:    time.Now,
		Timestep: DefaultTimeStep,

		tx: tx,
	}, nil
}

// Close terminates the session.
func (c *Card) Close() error {
	if c.tx != nil {
		if err := c.tx.EndTransaction(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Card) Select() error {
	_, err := c.Card.Select(iso.AidFeitianOTP)
	return err
}

func checkName(name string) error {
	if len(name) < 4 {
		return ErrNameTooShort
	} else if len(name) > 64 {
		return ErrNameTooLong
	}
	return nil
}

func checkSlot(slot Slot) error {
	if slot != Slot1 && slot != Slot2 && slot != SlotDefault {
		return ErrInvalidSlot
	}

	return nil
}

func checkDigits(digits int) error {
	if digits != 6 && digits != 8 {
		return ErrInvalidDigits
	}

	return nil
}
