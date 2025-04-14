// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian

import "cunicu.li/go-iso7816"

// Swap swaps the two slot configurations.
func (c *Card) Swap() error {
	_, err := c.Send(&iso7816.CAPDU{
		Ins: insSwapSlot,
	})
	return err
}
