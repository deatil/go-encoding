package encoding

import (
    "github.com/deatil/go-encoding/base58"
)

// Base58
func (this Encoding) Base58Decode() Encoding {
    data := string(this.data)
    this.data, this.Error = base58.StdEncoding.DecodeString(data)

    return this
}

// 编码 Base58
func (this Encoding) Base58Encode() Encoding {
    data := base58.StdEncoding.EncodeToString(this.data)
    this.data = []byte(data)

    return this
}
