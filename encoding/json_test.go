package encoding

import (
    "testing"
)

func Test_JSON(t *testing.T) {
    assert := assertT(t)
    assertError := assertErrorT(t)

    name := "JSON"
    data := map[string]string{
        "key": "test-pass",
        "key2": "test-pass-JSON",
    }

    en := FromString("").JSONEncode(data)
    enStr := en.ToString()

    assertError(en.Error, name + " Encode error")

    var deData map[string]string
    de := FromString(enStr).JSONDecode(&deData)

    assertError(de.Error, name + " Decode error")

    assert(data["key"], deData["key"], name)
    assert(data["key2"], deData["key2"], name)
}

