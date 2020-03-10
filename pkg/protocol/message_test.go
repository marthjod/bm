package protocol

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
				Header: MagicHeaderValue,
			},
			expected: []byte{0xe9, 0xbe, 0xb4, 0xd9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			msg: &Message{
				Header:  MagicHeaderValue,
				Command: "test",
			},
			expected: []byte{0xe9, 0xbe, 0xb4, 0xd9, 0x74, 0x65, 0x73, 0x74, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			msg: &Message{
				Header:  MagicHeaderValue,
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

// func TestVarIntListMarshalBinary(t *testing.T) {
// 	var tests = []struct {
// 		list        VarIntList
// 		expected    []byte
// 		expectedErr error
// 	}{
// 		{
// 			list: VarIntList{
// 				Count:  0,
// 				Values: []int{},
// 			},
// 			expected: nil,
// 		},
// 		{
// 			list: VarIntList{
// 				Count:  1,
// 				Values: []int{23},
// 			},
// 			expected: []byte{1, 23},
// 		},
// 		{
// 			list: VarIntList{
// 				Count:  2,
// 				Values: []int{23, 24},
// 			},
// 			expected: []byte{2, 23, 24},
// 		},
// 		{
// 			list: VarIntList{
// 				Count:  128,
// 				Values: []int{1},
// 			},
// 			expected:    nil,
// 			expectedErr: ErrLengthCountMismatch,
// 		},
// 		{
// 			list: VarIntList{
// 				Count:  128,
// 				Values: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127},
// 			},
// 			expected: []byte{128, 1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127},
// 		},
// 	}

// 	for _, test := range tests {
// 		marshaled, err := test.list.MarshalBinary()
// 		if err != nil {
// 			if test.expectedErr != nil {
// 				if err != test.expectedErr {
// 					t.Errorf("errors do not match: expected %q, got %q", test.expectedErr, err)
// 					return
// 				}
// 				continue
// 			}
// 			t.Error(err)
// 			return
// 		}
// 		if !bytes.Equal(marshaled, test.expected) {
// 			t.Errorf("\nexpected: % x,\ngot:      % x", test.expected, marshaled)
// 		}
// 	}
// }

// // TODO
// func TestMarshalBinaryVersion(t *testing.T) {
// 	var tests = []struct {
// 		msg      *MessageVersion
// 		expected []byte
// 	}{
// 		{
// 			msg:      Version(),
// 			expected: []byte{0xe9, 0xbe, 0xb4, 0xd9, 0x76, 0x65, 0x72, 0x61, 0x63, 0x6b, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		},
// 	}

// 	for _, test := range tests {
// 		marshaled, err := test.msg.MarshalBinary()
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		if !bytes.Equal(marshaled, test.expected) {
// 			t.Errorf("\nexpected: % x,\ngot:      % x", test.expected, marshaled)
// 		}
// 	}
// }
