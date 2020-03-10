package message

import (
	"bytes"
	"testing"
)

func TestMarshalBinary(t *testing.T) {
	var tests = []struct {
		msg      *Message
		expected []byte
	}{
		{
			msg:      &Message{},
			expected: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			msg: &Message{
				Magic: KnownMagicValue,
			},
			expected: []byte{0xe9, 0xbe, 0xb4, 0xd9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			msg: &Message{
				Magic:   KnownMagicValue,
				Command: "test",
			},
			expected: []byte{0xe9, 0xbe, 0xb4, 0xd9, 0x74, 0x65, 0x73, 0x74, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			msg: &Message{
				Magic:   KnownMagicValue,
				Command: "verack",
			},
			expected: []byte{0xe9, 0xbe, 0xb4, 0xd9, 0x76, 0x65, 0x72, 0x61, 0x63, 0x6b, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			msg:      VerAck(),
			expected: []byte{0xe9, 0xbe, 0xb4, 0xd9, 0x76, 0x65, 0x72, 0x61, 0x63, 0x6b, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		marshaled, err := test.msg.MarshalBinary()
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(marshaled, test.expected) {
			t.Errorf("\nexpected: % x,\ngot:      % x", test.expected, marshaled)
		}
	}
}
