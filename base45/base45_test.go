package base45

import (
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

// rfc9285Encode is the reference encoder from RFC 9285, Section 3: a two-byte
// group is always three characters, a trailing single byte is two.
func rfc9285Encode(data []byte) string {
    out := make([]byte, 0, len(data)/2*3+2)
    i := 0
    for ; i+1 < len(data); i += 2 {
        n := int(data[i])*256 + int(data[i+1])
        out = append(out, encodeStd[n%45], encodeStd[(n/45)%45], encodeStd[n/(45*45)])
    }
    if i < len(data) {
        a := int(data[i])
        out = append(out, encodeStd[a%45], encodeStd[a/45])
    }
    return string(out)
}

// TestEncodeAllTwoBytePairs checks every 16-bit value against the RFC 9285
// reference and confirms the round-trip is lossless. A full pair must encode to
// three characters even when its value is below 45*45 (e.g. \x00\x00 -> "000");
// emitting only two loses a byte on decode.
func TestEncodeAllTwoBytePairs(t *testing.T) {
    for v := 0; v < 65536; v++ {
        in := []byte{byte(v >> 8), byte(v)}

        got := StdEncoding.EncodeToString(in)
        if want := rfc9285Encode(in); got != want {
            t.Fatalf("Encode(% x) = %q, want %q", in, got, want)
        }
        if len(got) != 3 {
            t.Fatalf("Encode(% x) = %q, want 3 chars", in, got)
        }

        back, err := StdEncoding.DecodeString(got)
        if err != nil {
            t.Fatalf("Decode(%q): %v", got, err)
        }
        if string(back) != string(in) {
            t.Fatalf("round-trip % x -> %q -> % x", in, got, back)
        }
    }
}

// TestEncodeShortFinalPair pins the reported cases where the final group is a
// small full pair, including inputs longer than one pair.
func TestEncodeShortFinalPair(t *testing.T) {
    cases := []struct {
        in  []byte
        out string
    }{
        {[]byte{0x00, 0x00}, "000"},
        {[]byte{0x00, 0x01}, "100"},
        {[]byte{0xff, 0xff, 0x00, 0x00}, "FGW000"},
        {[]byte{0x12, 0x34, 0x00, 0x2a}, "PD2.00"},
        {[]byte{0x00}, "00"}, // genuine single-byte tail stays two chars
    }
    for _, c := range cases {
        got := StdEncoding.EncodeToString(c.in)
        if got != c.out {
            t.Errorf("Encode(% x) = %q, want %q", c.in, got, c.out)
        }
        back, err := StdEncoding.DecodeString(got)
        if err != nil || string(back) != string(c.in) {
            t.Errorf("round-trip % x -> %q -> % x (err %v)", c.in, got, back, err)
        }
    }
}

// TestDecodeEmpty confirms decoding an empty string yields empty output rather
// than panicking (RFC 9285 decodes empty to empty).
func TestDecodeEmpty(t *testing.T) {
    got, err := StdEncoding.DecodeString("")
    if err != nil {
        t.Fatalf("Decode(\"\"): %v", err)
    }
    if len(got) != 0 {
        t.Fatalf("Decode(\"\") = % x, want empty", got)
    }
}
