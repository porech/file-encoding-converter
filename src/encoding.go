package main

import (
	"github.com/porech/file-encoding-converter/src/rds"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

type EncoderDecoder interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, error)
}

type NullEncoderDecoder struct{}

func NewNullEncoderDecoder() EncoderDecoder {
	return &NullEncoderDecoder{}
}

func (_ NullEncoderDecoder) Encode(str string) ([]byte, error) {
	return []byte(str), nil
}

func (_ NullEncoderDecoder) Decode(bytes []byte) (string, error) {
	return string(bytes), nil
}

type BuiltinEncoderDecoder struct {
	encoder *encoding.Encoder
	decoder *encoding.Decoder
}

func NewCharmapEncoderDecoder(charmap *charmap.Charmap) EncoderDecoder {
	return &BuiltinEncoderDecoder{
		encoder: charmap.NewEncoder(),
		decoder: charmap.NewDecoder(),
	}
}

func NewUnicodeEncoderDecoder(charmap encoding.Encoding) EncoderDecoder {
	return &BuiltinEncoderDecoder{
		encoder: charmap.NewEncoder(),
		decoder: charmap.NewDecoder(),
	}
}

func (c BuiltinEncoderDecoder) Encode(str string) ([]byte, error) {
	return c.encoder.Bytes([]byte(str))
}

func (c BuiltinEncoderDecoder) Decode(bytes []byte) (string, error) {
	return c.decoder.String(string(bytes))
}

type RDSEncoderDecoder struct{}

func NewRDSEncoderDecoder() EncoderDecoder {
	return &RDSEncoderDecoder{}
}

func (_ RDSEncoderDecoder) Encode(str string) ([]byte, error) {
	return rds.Encode(str), nil
}

func (_ RDSEncoderDecoder) Decode(bytes []byte) (string, error) {
	return rds.Decode(bytes), nil
}
