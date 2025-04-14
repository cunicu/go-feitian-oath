// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian

import iso "cunicu.li/go-iso7816"

// Reset deletes all OTP credentials.
func (c *Card) Reset() error {
	_, err := c.Send(&iso.CAPDU{
		Ins: insReset,
	})
	return err
}
