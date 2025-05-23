// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian

import (
	"errors"

	iso "cunicu.li/go-iso7816"
)

type Language byte

const (
	LangEnglish Language = 0
	LangFrench  Language = 2
)

var ErrUnsupportedLanguage = errors.New("unsupported language")

//nolint:gochecknoglobals
var languageCodes = map[Language][]byte{
	LangEnglish: {0x0D, 0x00, 0x28, 0x31, 0x00, 0x1E, 0x32, 0x00, 0x1F, 0x33, 0x00, 0x20, 0x34, 0x00, 0x21, 0x35, 0x00, 0x22, 0x36, 0x00, 0x23, 0x37, 0x00, 0x24, 0x38, 0x00, 0x25, 0x39, 0x00, 0x26, 0x30, 0x00, 0x27, 0x61, 0x00, 0x04, 0x62, 0x00, 0x05, 0x63, 0x00, 0x06, 0x64, 0x00, 0x07, 0x65, 0x00, 0x08, 0x66, 0x00, 0x09, 0x67, 0x00, 0x0A, 0x68, 0x00, 0x0B, 0x69, 0x00, 0x0C, 0x6A, 0x00, 0x0D, 0x6B, 0x00, 0x0E, 0x6C, 0x00, 0x0F, 0x6D, 0x00, 0x10, 0x6E, 0x00, 0x11, 0x6F, 0x00, 0x12, 0x70, 0x00, 0x13, 0x71, 0x00, 0x14, 0x72, 0x00, 0x15, 0x73, 0x00, 0x16, 0x74, 0x00, 0x17, 0x75, 0x00, 0x18, 0x76, 0x00, 0x19, 0x77, 0x00, 0x1A, 0x78, 0x00, 0x1B, 0x79, 0x00, 0x1C, 0x7A, 0x00, 0x1D, 0x41, 0x02, 0x04, 0x42, 0x02, 0x05, 0x43, 0x02, 0x06, 0x44, 0x02, 0x07, 0x45, 0x02, 0x08, 0x46, 0x02, 0x09, 0x47, 0x02, 0x0A, 0x48, 0x02, 0x0B, 0x49, 0x02, 0x0C, 0x4A, 0x02, 0x0D, 0x4B, 0x02, 0x0E, 0x4C, 0x02, 0x0F, 0x4D, 0x02, 0x10, 0x4E, 0x02, 0x11, 0x4F, 0x02, 0x12, 0x50, 0x02, 0x13, 0x51, 0x02, 0x14, 0x52, 0x02, 0x15, 0x53, 0x02, 0x16, 0x54, 0x02, 0x17, 0x55, 0x02, 0x18, 0x56, 0x02, 0x19, 0x57, 0x02, 0x1A, 0x58, 0x02, 0x1B, 0x59, 0x02, 0x1C, 0x5A, 0x02, 0x1D, 0x27, 0x00, 0x35, 0x2D, 0x00, 0x2D, 0x3D, 0x00, 0x2E, 0x5B, 0x00, 0x2F, 0x5D, 0x00, 0x30, 0x3B, 0x00, 0x33, 0x60, 0x00, 0x34, 0x5C, 0x00, 0x31, 0x2C, 0x00, 0x36, 0x2E, 0x00, 0x37, 0x2F, 0x00, 0x38, 0x20, 0x00, 0x2C, 0x7E, 0x02, 0x35, 0x5F, 0x02, 0x2D, 0x2B, 0x02, 0x2E, 0x7B, 0x02, 0x2F, 0x7D, 0x02, 0x30, 0x3A, 0x02, 0x33, 0x22, 0x02, 0x34, 0x7C, 0x02, 0x31, 0x3C, 0x02, 0x36, 0x3E, 0x02, 0x37, 0x3F, 0x02, 0x38, 0x21, 0x02, 0x1E, 0x40, 0x02, 0x1F, 0x23, 0x02, 0x20, 0x24, 0x02, 0x21, 0x25, 0x02, 0x22, 0x5E, 0x02, 0x23, 0x26, 0x02, 0x24, 0x2A, 0x02, 0x25, 0x28, 0x02, 0x26, 0x29, 0x02, 0x27},
	LangFrench:  {0x0D, 0x00, 0x28, 0x31, 0x02, 0x1E, 0x32, 0x02, 0x1F, 0x33, 0x02, 0x20, 0x34, 0x02, 0x21, 0x35, 0x02, 0x22, 0x36, 0x02, 0x23, 0x37, 0x02, 0x24, 0x38, 0x02, 0x25, 0x39, 0x02, 0x26, 0x30, 0x02, 0x27, 0x61, 0x00, 0x14, 0x62, 0x00, 0x05, 0x63, 0x00, 0x06, 0x64, 0x00, 0x07, 0x65, 0x00, 0x08, 0x66, 0x00, 0x09, 0x67, 0x00, 0x0A, 0x68, 0x00, 0x0B, 0x69, 0x00, 0x0C, 0x6A, 0x00, 0x0D, 0x6B, 0x00, 0x0E, 0x6C, 0x00, 0x0F, 0x6D, 0x00, 0x33, 0x6E, 0x00, 0x11, 0x6F, 0x00, 0x12, 0x70, 0x00, 0x13, 0x71, 0x00, 0x04, 0x72, 0x00, 0x15, 0x73, 0x00, 0x16, 0x74, 0x00, 0x17, 0x75, 0x00, 0x18, 0x76, 0x00, 0x19, 0x77, 0x00, 0x1D, 0x78, 0x00, 0x1B, 0x79, 0x00, 0x1C, 0x7A, 0x00, 0x1A, 0x41, 0x02, 0x14, 0x42, 0x02, 0x05, 0x43, 0x02, 0x06, 0x44, 0x02, 0x07, 0x45, 0x02, 0x08, 0x46, 0x02, 0x09, 0x47, 0x02, 0x0A, 0x48, 0x02, 0x0B, 0x49, 0x02, 0x0C, 0x4A, 0x02, 0x0D, 0x4B, 0x02, 0x0E, 0x4C, 0x02, 0x0F, 0x4D, 0x02, 0x33, 0x4E, 0x02, 0x11, 0x4F, 0x02, 0x12, 0x50, 0x02, 0x13, 0x51, 0x02, 0x04, 0x52, 0x02, 0x15, 0x53, 0x02, 0x16, 0x54, 0x02, 0x17, 0x55, 0x02, 0x18, 0x56, 0x02, 0x19, 0x57, 0x02, 0x1D, 0x58, 0x02, 0x1B, 0x59, 0x02, 0x1C, 0x5A, 0x02, 0x1A, 0x27, 0x40, 0x24, 0x2D, 0x00, 0x23, 0x3D, 0x00, 0x2E, 0x5B, 0x40, 0x22, 0x5D, 0x40, 0x2D, 0x3B, 0x00, 0x36, 0x60, 0x40, 0x24, 0x5C, 0x40, 0x25, 0x2C, 0x00, 0x10, 0x2E, 0x02, 0x36, 0x2F, 0x02, 0x37, 0x20, 0x00, 0x2C, 0x7E, 0x40, 0x1F, 0x5F, 0x00, 0x25, 0x2B, 0x02, 0x2E, 0x7B, 0x40, 0x21, 0x7D, 0x40, 0x2E, 0x3A, 0x00, 0x37, 0x22, 0x02, 0x34, 0x7C, 0x40, 0x23, 0x3F, 0x02, 0x10, 0x21, 0x00, 0x38, 0x40, 0x40, 0x27, 0x23, 0x40, 0x20, 0x24, 0x00, 0x30, 0x25, 0x02, 0x34, 0x5E, 0x40, 0x26, 0x26, 0x00, 0x1E, 0x2A, 0x00, 0x31, 0x28, 0x00, 0x22, 0x29, 0x00, 0x2D},
}

// SetLanguage sets English and French key value codes.
func (c *Card) SetLanguage(lang Language) error {
	codes, ok := languageCodes[lang]
	if !ok {
		return ErrUnsupportedLanguage
	}

	_, err := c.Send(&iso.CAPDU{
		Ins:  insLanguage,
		P1:   0x00,
		P2:   0x01,
		Data: codes,
	})

	return err
}

// Language returns the currently configured language.
func (c *Card) Language() (Language, error) {
	resp, err := c.Send(&iso.CAPDU{
		Ins:  insLanguage,
		Data: []byte{0x31},
	})
	if err != nil {
		return 0, err
	} else if len(resp) < 1 {
		return 0, iso.ErrWrongLength
	}

	if resp[0] == 2 {
		return LangFrench, nil
	}

	return LangEnglish, nil
}
