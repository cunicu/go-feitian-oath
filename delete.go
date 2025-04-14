// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian

// Delete removes the configuration from a slot.
func (c *Card) Delete(_ /*slot*/ Slot, _ /*name*/ string) error {
	return nil
}

//     try:
//         if slot not in list(DICT_SLOT.keys()):
//             result_dict = {RES_SW: '%X' % SW.ERR_CUSTOMER,
//              RES_INFO: 'slot must be in {}.'.format(list(DICT_SLOT.keys()))}
//             return
//         name = name.replace(' ', '')
//         if name is None or len(name) < 4 or len(name) > 64:
//             result_dict = {RES_SW: '%X' % SW.ERR_CUSTOMER,
//              RES_INFO: 'name error.'}
//             return
//         data = Tlv(TAG.OTP_NAME, bytes(name, encoding='utf-8'))
//         apdu = APDU.build(0, INS.OTP_DELETE, 0, slot, data)
//         result, sw = dev.sendApdu(apdu)
//         result_dict.update(SW.get_info(sw))
//     finally:
//         click.echo(str(result_dict))
