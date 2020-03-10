package protocol

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"net"
	"time"
)

const (
	// MagicHeaderValue see https://bitmessage.org/wiki/Protocol_specification#Message_structure.
	MagicHeaderValue = 0xe9beb4d9
	// ProtocolVersion is the supported BitMessage protocol version.
	ProtocolVersion = 3
	// VerAckString is the `verack` message command string.
	VerAckString = "verack"
	// VersionString is the `version` message command string.
	VersionString = "version"
	// UserAgent is the default user agent.
	UserAgent = "go-bitmessage/bm"
)

type VarIntList struct {
	Count  int
	Values []int
}

type VarString struct {
	Len    uint64
	String string
}

func (v *VarString) MarshalBinary() (data []byte, err error) {
	var buf = make([]byte, uint64(binary.Size(v.Len))+v.Len)

	// panics if buffer too small
	binary.PutUvarint(buf, v.Len)

	bw := bytes.NewBuffer(buf)
	if err := binary.Write(bw, binary.BigEndian, []byte(v.String)); err != nil {
		return nil, err
	}

	return bw.Bytes(), nil
}

// Message is a BitMessage message, see https://bitmessage.org/wiki/Protocol_specification.
type Message struct {
	Header   uint32
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
		uint32(m.Header),
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

// NetworkAddress represents a network address.
type NetworkAddress struct {
	net.TCPAddr
	Time     uint64
	Stream   uint32
	Services uint64
}

// func (n *NetworkAddress) MarshalBinary() (data []byte, err error) {

// }

type MessageVersion struct {
	Message
	Version       int32
	Services      uint64
	Timestamp     time.Time
	Receiver      NetworkAddress
	Sender        NetworkAddress
	Nonce         uint64
	UserAgent     string
	StreamNumbers VarIntList
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (m *MessageVersion) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	var cmd [12]byte
	copy(cmd[:], m.Command)

	ua := VarString{
		Len:    uint64(len(m.UserAgent)),
		String: m.UserAgent,
	}
	userAgent, err := ua.MarshalBinary()
	if err != nil {
		return nil, err
	}

	// sn := VarIntList{
	// 	Count:  0,
	// 	Values: []int{},
	// }
	// streamNumbers, err := sn.MarshalBinary()
	// if err != nil {
	// 	return nil, err
	// }

	var msg = []interface{}{
		uint32(m.Header),
		[12]byte(cmd),
		uint64(m.Services),
		m.Timestamp.UnixNano(),
		m.Receiver.IP.To16(),
		uint16(m.Receiver.Port),
		m.Sender.IP.To16(),
		uint16(m.Receiver.Port),
		nonce(),
		userAgent,
		// streamNumbers,
	}
	for _, v := range msg {
		if err := binary.Write(buf, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// VerAck returns a valid `verack` message.
func VerAck() *Message {
	return &Message{
		Header:  MagicHeaderValue,
		Command: VerAckString,
	}
}

// Version returns a valid `version` message.
func Version() *MessageVersion {
	return &MessageVersion{
		Message: Message{
			Header:  MagicHeaderValue,
			Command: VersionString,
		},
		Version:   ProtocolVersion,
		UserAgent: UserAgent,
		// TODO: local address
	}
}

func nonce() uint64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint64()
}
