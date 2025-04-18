## go-encoding

<p align="center">
<a href="https://pkg.go.dev/github.com/deatil/go-encoding" ><img src="https://pkg.go.dev/badge/deatil/go-encoding.svg" alt="Go Reference"></a>
<a href="https://codecov.io/gh/deatil/go-encoding" ><img src="https://codecov.io/gh/deatil/go-encoding/graph/badge.svg?token=SS2Z1IY0XL"/></a>
<img src="https://goreportcard.com/badge/github.com/deatil/go-encoding" />
</p>

中文 | [English](README.md)

### 项目介绍

*  常用的编码解码算法
*  算法包括: (Hex/Base32/Base36/Base45/Base58/Base62/Base64/Base85/Base91/Base92/Base100/MorseITU/JSON)


### 下载安装

~~~go
go get -u github.com/deatil/go-encoding
~~~


### 开始使用

~~~go
package main

import (
    "fmt"
    "github.com/deatil/go-encoding/encoding"
)

func main() {
    oldData := "useData"

    // Base64 编码
    base64Data := encoding.
        FromString(oldData).
        Base64Encode().
        ToString()
    fmt.Println("Base64 编码为：", base64Data)

    // Base64 解码
    base64DecodeData := encoding.
        FromString(base64Data).
        Base64Decode().
        ToString()
    fmt.Println("Base64 解码为：", base64DecodeData)
}
~~~


### 格式说明

~~~go
base64Data := encoding.
    FromString(oldData). // 输入数据
    Base64Encode().      // 编码方式/解码方式
    ToString()           // 输出数据
~~~


### 输入输出数据

*  输入数据:
`FromBytes(data []byte)`, `FromString(data string)`, `FromReader(reader io.Reader)`
*  输出数据:
`String() string`, `ToBytes() []byte`, `ToString() string`, `ToReader() io.Reader`


### 常用解码编码

*  编码方式:
`Base32Encode()`, `Base32RawEncode()`,  `Base32HexEncode()`,`Base32RawHexEncode()`,  `Base32EncodeWithEncoder(encoder string)`, `Base32RawEncodeWithEncoder(encoder string)`,
`Base45Encode()`,
`Base58Encode()`,
`Base62Encode()`,
`Base64Encode()`, `Base64URLEncode()`, `Base64RawEncode()`, `Base64RawURLEncode()`, `Base64SegmentEncode()`, `Base64EncodeWithEncoder(encoder string)`,
`Base85Encode()`,
`Base91Encode()`,
`Base100Encode()`,
`Basex2Encode()`, `Basex16Encode()`, `Basex62Encode()`, `BasexEncodeWithEncoder(encoder string)`,
`HexEncode()`,
`MorseITUEncode()`,
`SafeURLEncode()`,
`SerializeEncode()`,
`JSONEncode(data any)`, `JSONIteratorEncode(data any)`, `JSONIteratorIndentEncode(v any, prefix, indent string)`,
`GobEncode(data any)`

*  解码方式:
`Base32Decode()`, `Base32RawDecode()`,  `Base32HexDecode()`,`Base32RawHexDecode()`,  `Base32DecodeWithEncoder(encoder string)`, `Base32RawDecodeWithEncoder(encoder string)`,
`Base45Decode()`,
`Base58Decode()`,
`Base62Decode()`,
`Base64Decode()`, `Base64URLDecode()`, `Base64RawDecode()`, `Base64RawURLDecode()`, `Base64SegmentDecode(paddingAllowed ...bool)`, `Base64DecodeWithEncoder(encoder string)`,
`Base85Encode()`,
`Base91Decode()`,
`Base100Decode()`,
`Basex2Decode()`, `Basex16Decode()`, `Basex62Decode()`, `BasexDecodeWithEncoder(encoder string)`,
`HexDecode()`,
`MorseITUDecode()`,
`SafeURLDecode()`,
`SerializeDecode()`,
`JSONDecode(dst any)`, `JSONIteratorDecode(dst any)`,
`GobDecode(dst any)`


### 开源协议

*  本软件包遵循 `Apache2` 开源协议发布，在保留本软件包版权的情况下提供个人及商业免费使用。


### 版权

*  本软件包所属版权归 deatil(https://github.com/deatil) 所有。
