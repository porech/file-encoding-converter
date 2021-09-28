package rds

import (
	"reflect"
	"testing"
)

func Test_Encode(t *testing.T) {
	var testCases = map[string][]byte{
		"Ciao":     {0x43, 0x69, 0x61, 0x6f},
		"Felicità": {0x46, 0x65, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x81},
		"KE$HA":    {0x4b, 0x45, 0xab, 0x48, 0x41},
		"I♡You":    {0x49, 0x59, 0x6f, 0x75},
	}

	for source, expected := range testCases {
		res := Encode(source)
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("wrong conversion for '%s': expected '%v', got '%v'", source, expected, res)
		}
	}
}

func Test_Decode(t *testing.T) {
	var testCases = map[string][]byte{
		"Ciao":     {0x43, 0x69, 0x61, 0x6f},
		"Felicità": {0x46, 0x65, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x81},
		"KE$HA":    {0x4b, 0x45, 0xab, 0x48, 0x41},
		"IYou":     {0x49, 0x59, 0x6f, 0x75},
	}

	for expected, source := range testCases {
		res := Decode(source)
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("wrong conversion for '%s': expected '%v', got '%v'", source, expected, res)
		}
	}
}
