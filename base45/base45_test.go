package base45

import (
    "bytes"
    "testing"
)

type testPair struct {
    decoded string
    encoded string
}

type malformedTestPair struct {
    encoded string
    error error
}

var pairs = []testPair{
    {
        decoded: "AB",
        encoded: "BB8",
    },
    {
        decoded: "Hello!!",
        encoded: "%69 VD92EX0",
    },
    {
        decoded: "base-45",
        encoded: "UJCLQE7W581",
    },
    {
        decoded: "ietf!",
        encoded: "QED8WEX0",
    },
    {
        decoded: "Lorem ipsum dolor sit amet",
        encoded: "$T9ZKE ZD$ED$QE ZDGVC*VDBJEPQESUEBEC7$C",
    },
}

var malformedPairs = []malformedTestPair{
    {
        encoded: "A",
        error:   InvalidLengthError{length: 1, mod: 1},
    },
    {
        encoded: "FAILEd",
        error:   InvalidCharacterError{char: 'd', position: 5},
    },
    {
        encoded: "GGW",
        error:   IllegalBase45ByteError{position: 0},
    },
    {
        encoded: ":::",
        error:   IllegalBase45ByteError{position: 0},
    },
    {
        encoded: "ABC:::",
        error:   IllegalBase45ByteError{position: 1},
    },
}

func testEqual(t *testing.T, msg string, args ...any) bool {
    t.Helper()
    if args[len(args)-2] != args[len(args)-1] {
        t.Errorf(msg, args...)
        return false
    }
    return true
}

func TestEncode(t *testing.T) {
    for _, p := range pairs {
        got := StdEncoding.EncodeToString([]byte(p.decoded))
        testEqual(t, "Encode(%q) = %q, want %q", p.decoded, got, p.encoded)
    }
}

func TestDecode(t *testing.T) {
    for _, p := range pairs {
        got, err := StdEncoding.DecodeString(p.encoded)
        if err != nil {
            t.Failed()
        }

        testEqual(t, "Decode(%q) = %q, want %q", p.encoded, string(got), p.decoded)
    }
}

func TestDecodeMalformed(t *testing.T) {
    for _, p := range malformedPairs {
        _, err := StdEncoding.DecodeString(p.encoded)
        if err == nil {
            t.Failed()
        }
        switch err := err.(type) {
        case InvalidLengthError:
            testEqual(t, "Error(%s) = %s, want %s", p.error, err, p.error)
        case InvalidCharacterError:
            testEqual(t, "Error(%s) = %s, want %s", p.error, err, p.error)
        case IllegalBase45ByteError:
            testEqual(t, "Error(%s) = %s, want %s", p.error, err, p.error)
        }
    }
}

// TestEncodeEvenLengthTrailingZeroChar is a regression test for the bug where
// Encode dropped the 3rd base45 digit (e) of a 2-byte trailing pair whenever
// e == 0, emitting 2 chars instead of 3 and breaking the round-trip
// (Decode(Encode([00 10])) returned [10], silently dropping the 0x00 byte).
//
// Per RFC 9285 §4, an even-length input MUST produce an output whose length
// is divisible by 3 (3 chars per pair); 2 chars are emitted ONLY for a
// trailing single byte (odd input length).
func TestEncodeEvenLengthTrailingZeroChar(t *testing.T) {
    cases := []testPair{
        // 2-byte pairs whose 3rd base45 digit (e) is 0 — previously dropped.
        {decoded: "\x00\x10", encoded: "G00"},
        {decoded: "\x00\x00", encoded: "000"},
        {decoded: "\x00\xff", encoded: "U50"},
        // 4-byte inputs whose last pair has e == 0.
        {decoded: "\x01\x00\x00\x10", encoded: "V50G00"},
        {decoded: "\x00\x00\x00\x00", encoded: "000000"},
        // Controls: last pair has e != 0 (already correct).
        {decoded: "AB", encoded: "BB8"},
        // Controls: odd-length inputs (trailing single byte) — 2 chars correct.
        {decoded: "!", encoded: "X0"},
        {decoded: "\x00", encoded: "00"},
    }
    for _, p := range cases {
        got := StdEncoding.EncodeToString([]byte(p.decoded))
        testEqual(t, "Encode(%q) = %q, want %q", p.decoded, got, p.encoded)
    }
}

// TestEncodeRoundTripEvenLengthTrailingZeroChar verifies that Decode(Encode(x))
// == x for inputs that previously broke the round-trip.
func TestEncodeRoundTripEvenLengthTrailingZeroChar(t *testing.T) {
    inputs := [][]byte{
        {0x00, 0x10},
        {0x00, 0x00},
        {0x01, 0x00, 0x00, 0x10},
        {0x00, 0x00, 0x00, 0x00},
    }
    for _, in := range inputs {
        enc := StdEncoding.EncodeToString(in)
        dec, err := StdEncoding.DecodeString(enc)
        if err != nil {
            t.Errorf("Decode(Encode(% x)) returned error: %v", in, err)
            continue
        }
        if !bytes.Equal(dec, in) {
            t.Errorf("Decode(Encode(% x)) = % x, want % x", in, dec, in)
        }
    }
}

// TestEncodeRFC9285Vectors re-asserts the RFC 9285 test vectors.
func TestEncodeRFC9285Vectors(t *testing.T) {
    cases := []testPair{
        {decoded: "AB", encoded: "BB8"},
        {decoded: "Hello!!", encoded: "%69 VD92EX0"},
        {decoded: "ietf!", encoded: "QED8WEX0"},
    }
    for _, p := range cases {
        got := StdEncoding.EncodeToString([]byte(p.decoded))
        testEqual(t, "Encode(%q) = %q, want %q", p.decoded, got, p.encoded)
    }
}
