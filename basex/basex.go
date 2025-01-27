package basex

import (
    "fmt"
    "bytes"
    "errors"
    "unsafe"
    "reflect"
    "strconv"
    "math/big"
)

// 常用 key
const (
    Base2Key         = "01"
    Base16Key        = "0123456789ABCDEF"
    Base16InvalidKey = "0123456789abcdef"
    Base32Key        = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
    Base58Key        = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
    Base62Key        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    Base62_2Key      = "vPh7zZwA2LyU4bGq5tcVfIMxJi6XaSoK9CNp0OWljYTHQ8REnmu31BrdgeDkFs"
    Base62InvalidKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
    // 编码类型
    Base2Encoding            = NewEncoding(Base2Key)
    Base16Encoding           = NewEncoding(Base16Key)
    Base16InvalidKeyEncoding = NewEncoding(Base16InvalidKey)
    Base32Encoding           = NewEncoding(Base32Key)
    Base58Encoding           = NewEncoding(Base58Key)
    Base62Encoding           = NewEncoding(Base62Key)
    Base62_2Encoding         = NewEncoding(Base62_2Key)
    Base62InvalidEncoding    = NewEncoding(Base62InvalidKey)
)

// Basex
type Encoding struct {
    base        *big.Int
    alphabet    []rune
    alphabetMap map[rune]int

    Error       error
}

// 构造函数
// Example alphabets:
//   - base2: 01
//   - base16: 0123456789abcdef
//   - base32: 0123456789ABCDEFGHJKMNPQRSTVWXYZ
//   - base58: 123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
//   - base62: 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
func NewEncoding(alphabet string) *Encoding {
    runes := []rune(alphabet)
    runeMap := make(map[rune]int)

    enc := &Encoding{}

    for i := 0; i < len(runes); i++ {
        if _, ok := runeMap[runes[i]]; ok {
            enc.Error = errors.New("go-encoding/basex: Ambiguous alphabet.")
            return enc
        }

        runeMap[runes[i]] = i
    }

    enc.base = big.NewInt(int64(len(runes)))
    enc.alphabet = runes
    enc.alphabetMap = runeMap

    return enc
}

// 编码
func (enc *Encoding) Encode(source []byte) []byte {
    if len(source) == 0 {
        return nil
    }

    var (
        res bytes.Buffer
        k   = 0
    )
    for ; source[k] == 0 && k < len(source)-1; k++ {
        res.WriteRune(enc.alphabet[0])
    }

    var (
        mod big.Int
        sourceInt = new(big.Int).SetBytes(source)
    )

    for sourceInt.Uint64() > 0 {
        sourceInt.DivMod(sourceInt, enc.base, &mod)
        res.WriteRune(enc.alphabet[mod.Uint64()])
    }

    var (
        buf = res.Bytes()
        j   = len(buf) - 1
    )

    for k < j {
        buf[k], buf[j] = buf[j], buf[k]
        k++
        j--
    }

    return buf
}

// EncodeToString returns the basex encoding of src.
func (enc *Encoding) EncodeToString(src []byte) string {
    buf := enc.Encode(src)
    return string(buf)
}

// 解码
func (enc *Encoding) Decode(source []byte) ([]byte, error) {
    if len(source) == 0 {
        return nil, nil
    }

    var (
        data = []rune(string(source))
        dest = big.NewInt(0)
    )

    for i := 0; i < len(data); i++ {
        value, ok := enc.alphabetMap[data[i]]
        if !ok {
            return nil, errors.New("go-encoding/basex: non Base Character")
        }

        dest.Mul(dest, enc.base)
        if value > 0 {
            dest.Add(dest, big.NewInt(int64(value)))
        }
    }

    k := 0
    for ; data[k] == enc.alphabet[0] && k < len(data)-1; k++ {
    }

    buf := dest.Bytes()
    res := make([]byte, k, k+len(buf))

    return append(res, buf...), nil
}

// DecodeString returns the bytes represented by the basex string s.
func (enc *Encoding) DecodeString(s string) ([]byte, error) {
    sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
    bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
    return enc.Decode(*(*[]byte)(unsafe.Pointer(&bh)))
}

// 补码
func (enc *Encoding) padding(s string, minlen int) string {
    if len(s) >= minlen {
        return s
    }

    format := fmt.Sprint(`%0`, strconv.Itoa(minlen), "s")
    return fmt.Sprintf(format, s)
}
