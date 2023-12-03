package mtom

import (
	"encoding/base64"
	"errors"
)

var ErrInvalidInputData = errors.New("mtom: invalid input data")

func FromBytes(src []byte) (dst []byte) {
	base64.StdEncoding.Encode(dst, src)
	return
}

func ToBytes(src []byte) (dst []byte, err error) {
	if _, err := base64.StdEncoding.Decode(dst, src); err != nil {
		return nil, ErrInvalidInputData
	}

	return dst, nil
}
