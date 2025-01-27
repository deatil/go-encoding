package encoding

import (
    "github.com/deatil/go-encoding/basex"
)

// Basex2
func (this Encoding) Basex2Decode() Encoding {
    data := string(this.data)
    this.data, this.Error = basex.Base2Encoding.DecodeString(data)

    return this
}

// 编码 Base2
func (this Encoding) Basex2Encode() Encoding {
    data := basex.Base2Encoding.EncodeToString(this.data)
    this.data = []byte(data)

    return this
}

// ====================

// Basex16
func (this Encoding) Basex16Decode() Encoding {
    data := string(this.data)
    this.data, this.Error = basex.Base16Encoding.DecodeString(data)

    return this
}

// 编码 Base16
func (this Encoding) Basex16Encode() Encoding {
    data := basex.Base16Encoding.EncodeToString(this.data)
    this.data = []byte(data)

    return this
}

// ====================

// Basex62
func (this Encoding) Basex62Decode() Encoding {
    data := string(this.data)
    this.data, this.Error = basex.Base62Encoding.DecodeString(data)

    return this
}

// 编码 Basex62
func (this Encoding) Basex62Encode() Encoding {
    data := basex.Base62Encoding.EncodeToString(this.data)
    this.data = []byte(data)

    return this
}

// ====================

// BasexDecodeWithEncoder
func (this Encoding) BasexDecodeWithEncoder(encoder string) Encoding {
    data := string(this.data)
    this.data, this.Error = basex.NewEncoding(encoder).DecodeString(data)

    return this
}

// BasexEncodeWithEncoder
func (this Encoding) BasexEncodeWithEncoder(encoder string) Encoding {
    data := basex.NewEncoding(encoder).EncodeToString(this.data)
    this.data = []byte(data)

    return this
}
