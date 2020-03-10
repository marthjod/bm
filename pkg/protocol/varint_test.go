package protocol

import (
	"bytes"
	"testing"
)

func TestVarIntMarshalBinary(t *testing.T) {
	var tests = []struct {
		varint      VarInt
		expected    []byte
		expectedLen int
	}{
		{
			varint:      VarInt{},
			expected:    []byte{0},
			expectedLen: 1,
		},
		{
			varint:      VarInt{Value: 1},
			expected:    []byte{1},
			expectedLen: 1,
		},
		{
			varint:      VarInt{Value: uint8Limit - 1},
			expected:    []byte{0xfc},
			expectedLen: 1,
		},
		{
			varint:      VarInt{Value: uint8Limit},
			expected:    []byte{0xfd, 0, 0xfd},
			expectedLen: 3,
		},
		{
			varint:      VarInt{Value: uint16Limit},
			expected:    []byte{0xfd, 0xff, 0xff},
			expectedLen: 3,
		},
		{
			varint:      VarInt{Value: uint32Limit},
			expected:    []byte{0xfe, 0xff, 0xff, 0xff, 0xff},
			expectedLen: 5,
		},
		{
			varint:      VarInt{Value: uint32Limit + 1},
			expected:    []byte{0xff, 0, 0, 0, 1, 0, 0, 0, 0},
			expectedLen: 9,
		},
	}
	for _, test := range tests {
		marshaled, err := test.varint.MarshalBinary()
		if err != nil {
			t.Error(err)
		}
		if len(marshaled) != test.expectedLen {
			t.Errorf("expected length %d, got %d", test.expectedLen, len(marshaled))
		}
		if !bytes.Equal(marshaled, test.expected) {
			t.Errorf("\nexpected: % x,\ngot:      % x", test.expected, marshaled)
		}
	}
}
