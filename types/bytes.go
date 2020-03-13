package types

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

// Bytes represents a yaml marshalable []byte
type Bytes []byte

func parseBytes(s string, parser func(string) ([]byte, error)) (Bytes, error) {
	r, err := parser(s)
	if err != nil {
		return nil, err
	}
	return Bytes(r), nil
}

// ParseBytesHex parses an hex string
func ParseBytesHex(h string) (Bytes, error) { return parseBytes(h, hex.DecodeString) }

// ParseBytesBase64 parses a base64 string
func ParseBytesBase64(b64 string) (Bytes, error) {
	return parseBytes(b64, base64.RawStdEncoding.DecodeString)
}

// Hex returns an hex string
func (b Bytes) Hex() string { return hex.EncodeToString(b) }

// Base64 returns a base64 string
func (b Bytes) Base64() string { return base64.RawStdEncoding.EncodeToString(b) }

// MarshalJSON implements json.Marshaler
func (b Bytes) MarshalJSON() ([]byte, error) { return json.Marshal(b.Hex()) }

// UnmarshalJSON implements json.Unmarshaler
func (b *Bytes) UnmarshalJSON(bb []byte) error {
	var s string
	err := json.Unmarshal(bb, &s)
	if err != nil {
		return err
	}
	tb, err := ParseBytesHex(s)
	if err != nil {
		return err
	}
	*b = tb
	return nil
}

func (b *Bytes) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var t string
	if err := unmarshal(&t); err != nil {
		return err
	}
	r, err := ParseBytesHex(t)
	if err != nil {
		return err
	}
	*b = Bytes(r)
	return nil
}

func (b Bytes) MarshalYAML() (interface{}, error) { return b.Hex(), nil }
