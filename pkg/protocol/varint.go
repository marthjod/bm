package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	uint8Limit  = 0xfd
	uint16Limit = 0xffff
	uint32Limit = 0xffffffff

	uint16Marker = 0xfd
	uint32Marker = 0xfe
	uint64Marker = 0xff
)

// VarInt represents a var_int.
type VarInt struct {
	Value int
}

// MarshalBinary implements binary.Marshaler.
func (v VarInt) MarshalBinary() (data []byte, err error) {
	// if v.Value == 0 {
	// 	return []byte{}, nil
	// }

	buf := new(bytes.Buffer)
	var out []interface{}

	if v.Value < uint8Limit {
		out = []interface{}{
			uint8(v.Value),
		}
	} else if v.Value <= uint16Limit {
		out = []interface{}{
			uint8(uint16Marker),
			uint16(v.Value),
		}
	} else if v.Value <= uint32Limit {
		out = []interface{}{
			uint8(uint32Marker),
			uint32(v.Value),
		}
	} else {
		out = []interface{}{
			uint8(uint64Marker),
			uint64(v.Value),
		}
	}
	for _, o := range out {
		if err := binary.Write(buf, binary.BigEndian, o); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
