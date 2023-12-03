package mtom

import (
	"encoding/base64"
	"errors"
)

var ErrInvalidInputData = errors.New("mtom: invalid input data")

func FromBytes(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func ToBytes(src []byte) ([]byte, error) {
	dst, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		return nil, ErrInvalidInputData
	}

	return dst, nil
}
