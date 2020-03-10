package message

import (
	"bytes"
	"encoding/binary"
)

const (
	// KnownMagicValue see https://bitmessage.org/wiki/Protocol_specification#Message_structure.
	KnownMagicValue = 0xE9BEB4D9
	// Version is the supported BitMessage protocol version.
	Version = 3
	// VerAckString is the `verack` message command string.
	VerAckString = "verack"
)

// type VarInt struct {
// 	Int int
// }

// type VarIntList struct {
// 	Count  VarInt
// 	Values []VarInt
// }

// type VarString struct {
// 	Len    int
// 	String string
// }

// Message is a BitMessage message, see https://bitmessage.org/wiki/Protocol_specification.
type Message struct {
	Magic    uint32
	Command  string
	Len      uint32
	Checksum uint32
	Payload  []byte
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (m *Message) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	var cmd [12]byte
	copy(cmd[:], m.Command)

	var msg = []interface{}{
		uint32(m.Magic),
		[12]byte(cmd),
		uint32(m.Len),
		uint32(m.Checksum),
		[]byte(m.Payload),
	}
	for _, v := range msg {
		if err := binary.Write(buf, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// type NetworkAddress struct {
// 	Time     uint64
// 	Stream   uint32
// 	Services uint64
// 	IPv6     net.IP
// 	Port     uint16
// }

// type MsgVersion struct {
// 	Version       int32
// 	Services      uint64
// 	Timestamp     time.Time
// 	Receiver      NetworkAddress
// 	Sender        NetworkAddress
// 	Nonce         uint64
// 	UserAgent     string
// 	StreamNumbers VarIntList
// }

// VerAck returns a valid `verack` message.
func VerAck() *Message {
	return &Message{
		Magic:   KnownMagicValue,
		Command: VerAckString,
	}
}
