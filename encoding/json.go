package encoding

import (
    "encoding/json"
)

// JSON Encode
func (this Encoding) JSONEncode(data any) Encoding {
    this.data, this.Error = json.Marshal(data)

    return this
}

// JSON Indent Encode
func (this Encoding) JSONIndentEncode(v any, prefix, indent string) Encoding {
    this.data, this.Error = json.MarshalIndent(v, prefix, indent)

    return this
}

// JSON Decode
func (this Encoding) JSONDecode(dst any) Encoding {
    this.Error = json.Unmarshal(this.data, dst)

    return this
}
