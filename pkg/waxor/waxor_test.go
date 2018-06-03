package waxor

import "testing"

func TestEncodeXor(t *testing.T) {
	encoded := EncodeXor("asdf")
	valid := "{xor}Piw7OQ=="
	if encoded != valid {
		t.Errorf("Encoded value was incorrect, got: %s, want: %s", encoded, valid)
	}
}

func TestDecodeXor(t *testing.T) {
	decoded := DecodeXor("{xor}Piw7OQ==")
	valid := "asdf"
	if decoded != valid {
		t.Errorf("Decoded value was incorrect, got: %s, want: %s", decoded, valid)
	}
}
