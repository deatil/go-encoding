package encoding

import (
    "github.com/deatil/go-encoding/morse"
)

// MorseITU
func (this Encoding) MorseITUDecode() Encoding {
    data, err := morse.ITUEncoding.DecodeString(string(this.data))

    this.data = data
    this.Error = err

    return this
}

// 编码 MorseITU
func (this Encoding) MorseITUEncode() Encoding {
    data := morse.ITUEncoding.EncodeToString(this.data)
    this.data = []byte(data)

    return this
}
