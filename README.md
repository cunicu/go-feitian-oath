<!--
SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
SPDX-License-Identifier: Apache-2.0
-->

# go-feitian-oath: Template repository for Go packages in the cunicu organization

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/cunicu/go-feitian-oath/test.yaml?style=flat-square)](https://github.com/cunicu/go-feitian-oath/actions)
[![goreportcard](https://goreportcard.com/badge/github.com/cunicu/go-feitian-oath?style=flat-square)](https://goreportcard.com/report/github.com/cunicu/go-feitian-oath)
[![Codecov branch](https://img.shields.io/codecov/c/github/cunicu/go-feitian-oath/main?style=flat-square&token=6XoWouQg6K)](https://app.codecov.io/gh/cunicu/go-feitian-oath/tree/main)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue?style=flat-square)](https://github.com/cunicu/go-feitian-oath/blob/main/LICENSES/Apache-2.0.txt)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cunicu/go-feitian-oath?style=flat-square)
[![Go Reference](https://pkg.go.dev/badge/github.com/cunicu/go-feitian-oath.svg)](https://pkg.go.dev/github.com/cunicu/go-feitian-oath)

The package `go-feitian-oath` implements the protocol to manage HOTP, TOTP, challenge/response and static password credentials on FEITIAN FIDO keys.
It uses PC/SC over CCID interface to communicate with the smart card.

## Features

- Calculation of
  - Time-based One-time Passwords (TOTP)
  - Hash-based One-time Passwords (HOTP)
  - Challenge Response (HMAC)
  - Static passwords
- Slot managment
  - Set default
  - Swap
- Credential management
  - Put
  - List
  - Delete
- Factory reset of applet

## Tested devices

- [FEITIAN ePass FIDO NFC K9Plus](https://www.ftsafe.com/Products/FIDO/NFC)
  - COS version 3301

## Limitations

The OTP applet of the FEITIAN keys has several limitations in comparison to its competition:

- It can only store 2 credentials
  - Each of those two credentials can of type TOTP, HOTP, challenge/response or static password
- It only supports SHA1 and SHA256 hash algorithms
- Credentials can not be protected with a PIN code
- Initial counter values for HOTP credentials can not be set

**Note:** The FEITIAN OTP applet show similarities to [Yubico's `ykneo-oath` applet](https://github.com/Yubico/ykneo-oath) when it was still open source.

## Contact

Please have a look at the contact page: [cunicu.li/docs/contact](https://cunicu.li/docs/contact).

## License

go-feitian-oath is licensed under the [Apache 2.0](./LICENSE) license.
