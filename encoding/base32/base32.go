package base32

import (
	"encoding/base32"
	"errors"
)

// drop the following letters:
// i - can be confused with 1,j
// l - can be confused with 1
// o - can be confused with 0
// u - can be confused with v
const encodeCompact = "0123456789abcdefghjkmnpqrstvwxyz"

// CompactEncoding is a lowercase base32 encoding using Crockford's alphabet
var CompactEncoding = base32.NewEncoding(encodeCompact).WithPadding(base32.NoPadding)

// EncodeToString encodes src using CompactEncoding
func EncodeToString(src []byte) string {
	return CompactEncoding.EncodeToString(src)
}

func DecodeString(src string) ([]byte, error) {
	dst, err := CompactEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}
	if CompactEncoding.EncodeToString(dst) != src {
		return dst, errors.New("uncanonical base32 input")
	}
	return dst, nil
}
