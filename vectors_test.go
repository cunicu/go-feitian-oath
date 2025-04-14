// SPDX-FileCopyrightText: 2018 Joern Barthel <joern.barthel@kreuzwerker.de>
// SPDX-FileCopyrightText: 2024-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package feitian_test

import (
	"encoding/hex"
	"time"

	"cunicu.li/go-feitian-oath"
)

type vector struct {
	Algorithm feitian.Algorithm
	Digits    int
	Secret    []byte
	Name      string
	Kind      feitian.Kind
	Time      time.Time
	Counter   uint32
	Code      string
	Challenge []byte
	Hash      []byte
}

func fromString(s string) []byte {
	return []byte(s)
}

func fromHex(s string) []byte {
	h, err := hex.DecodeString(s)
	if err != nil {
		panic("failed to parse hex: " + err.Error())
	}
	return h
}

//nolint:gochecknoglobals
var (
	testChallenge = fromHex("53656420757420706572737069")

	// See: https://www.rfc-editor.org/errata/eid2866
	testSecretSHA1   = fromString("12345678901234567890")
	testSecretSHA256 = fromString("12345678901234567890123456789012")

	vectorsTOTP = []vector{
		// RFC 6238 Appendix B - Test Vectors
		// See: https://datatracker.ietf.org/doc/html/rfc6238#appendix-B
		{Name: "rfc6238-test-01", Algorithm: feitian.SHA1, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA1, Time: time.Unix(59, 0), Code: "94287082"},
		{Name: "rfc6238-test-02", Algorithm: feitian.SHA256, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA256, Time: time.Unix(59, 0), Code: "46119246"},
		// {Name: "rfc6238-test-03", Alg: feitian.SHA512, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA512, Time: time.Unix(59, 0), Code: "90693936"},
		{Name: "rfc6238-test-04", Algorithm: feitian.SHA1, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA1, Time: time.Unix(1111111109, 0), Code: "07081804"},
		{Name: "rfc6238-test-05", Algorithm: feitian.SHA256, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA256, Time: time.Unix(1111111109, 0), Code: "68084774"},
		// {Name: "rfc6238-test-06", Alg: feitian.SHA512, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA512, Time: time.Unix(1111111109, 0), Code: "25091201"},
		{Name: "rfc6238-test-07", Algorithm: feitian.SHA1, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA1, Time: time.Unix(1111111111, 0), Code: "14050471"},
		{Name: "rfc6238-test-08", Algorithm: feitian.SHA256, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA256, Time: time.Unix(1111111111, 0), Code: "67062674"},
		// {Name: "rfc6238-test-09", Alg: feitian.SHA512, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA512, Time: time.Unix(1111111111, 0), Code: "99943326"},
		{Name: "rfc6238-test-10", Algorithm: feitian.SHA1, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA1, Time: time.Unix(1234567890, 0), Code: "89005924"},
		{Name: "rfc6238-test-11", Algorithm: feitian.SHA256, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA256, Time: time.Unix(1234567890, 0), Code: "91819424"},
		// {Name: "rfc6238-test-12", Alg: feitian.SHA512, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA512, Time: time.Unix(1234567890, 0), Code: "93441116"},
		{Name: "rfc6238-test-13", Algorithm: feitian.SHA1, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA1, Time: time.Unix(2000000000, 0), Code: "69279037"},
		{Name: "rfc6238-test-14", Algorithm: feitian.SHA256, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA256, Time: time.Unix(2000000000, 0), Code: "90698825"},
		// {Name: "rfc6238-test-15", Alg: feitian.SHA512, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA512, Time: time.Unix(2000000000, 0), Code: "38618901"},
		{Name: "rfc6238-test-16", Algorithm: feitian.SHA1, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA1, Time: time.Unix(20000000000, 0), Code: "65353130"},
		{Name: "rfc6238-test-17", Algorithm: feitian.SHA256, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA256, Time: time.Unix(20000000000, 0), Code: "77737706"},
		// {Name: "rfc6238-test-18", Alg: feitian.SHA512, Kind: feitian.TOTP, Digits: 8, Secret: testSecretSHA512, Time: time.Unix(20000000000, 0), Code: "47863826"},
	}

	vectorsHOTP = []vector{
		// RFC 4226 Appendix D - HOTP Algorithm: Test Values
		// See: https://datatracker.ietf.org/doc/html/rfc4226#page-32
		{Name: "rfc4226-test-00", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 0, Code: "755224", Hash: fromHex("cc93cf18508d94934c64b65d8ba7667fb7cde4b0")},
		{Name: "rfc4226-test-01", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 1, Code: "287082", Hash: fromHex("75a48a19d4cbe100644e8ac1397eea747a2d33ab")},
		{Name: "rfc4226-test-02", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 2, Code: "359152", Hash: fromHex("0bacb7fa082fef30782211938bc1c5e70416ff44")},
		{Name: "rfc4226-test-03", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 3, Code: "969429", Hash: fromHex("66c28227d03a2d5529262ff016a1e6ef76557ece")},
		{Name: "rfc4226-test-04", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 4, Code: "338314", Hash: fromHex("a904c900a64b35909874b33e61c5938a8e15ed1c")},
		{Name: "rfc4226-test-05", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 5, Code: "254676", Hash: fromHex("a37e783d7b7233c083d4f62926c7a25f238d0316")},
		{Name: "rfc4226-test-06", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 6, Code: "287922", Hash: fromHex("bc9cd28561042c83f219324d3c607256c03272ae")},
		{Name: "rfc4226-test-07", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 7, Code: "162583", Hash: fromHex("a4fb960c0bc06e1eabb804e5b397cdc4b45596fa")},
		{Name: "rfc4226-test-08", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 8, Code: "399871", Hash: fromHex("1b3c89f65e6c9e883012052823443f048b4332db")},
		{Name: "rfc4226-test-09", Algorithm: feitian.SHA1, Kind: feitian.HOTP, Digits: 6, Secret: testSecretSHA1, Counter: 9, Code: "520489", Hash: fromHex("1637409809a679dc698207310c8c7fc07290d9e5")},
	}

	vectorsChalResp = []vector{
		{Name: "chalresp-test-01", Algorithm: feitian.SHA1, Kind: feitian.ChallengeResponse, Digits: 6, Secret: testSecretSHA1, Challenge: testChallenge, Hash: fromHex("7aa340360e3c25e82c8d6a6c1fe8d397b7887177")},
		{Name: "chalresp-test-02", Algorithm: feitian.SHA256, Kind: feitian.ChallengeResponse, Digits: 6, Secret: testSecretSHA256, Challenge: testChallenge, Hash: fromHex("b09821d85415a0dc8fccef529f76011b28a570363506d15cee11de8945522df8")},
	}

	vectors = map[string][]vector{
		"TOTP":     vectorsTOTP,
		"HOTP":     vectorsHOTP,
		"ChalResp": vectorsChalResp,
	}
)
