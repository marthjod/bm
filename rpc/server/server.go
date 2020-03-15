package server

import (
	"context"
	"net"
	"time"

	"github.com/marthjod/bm/pkg/nonce"
	"github.com/marthjod/bm/rpc/proto"
)

// TODO set during build time
const (
	Version   = 3
	UserAgent = "bm/0.1-alpha"
)

type Server struct {
	proto.VersionerServer
	IPAddress net.IP
	Port      int32
}

func New(port int32) (*Server, error) {
	s := &Server{
		Port: port,
	}
	ip, err := getLocalIPAddress()
	if err != nil {
		return nil, err
	}
	s.IPAddress = ip
	return s, nil
}

// Version returns the server's version.
func (s *Server) Version(ctx context.Context, req *proto.VersionRequest) (*proto.VersionResponse, error) {
	return &proto.VersionResponse{
		Version: &proto.Version{
			Version:   Version,
			Timestamp: time.Now().Unix(),
			Nonce:     nonce.New(),
			UserAgent: UserAgent,
			Receiver: &proto.NetworkAddress{
				Time:      uint64(time.Now().Unix()),
				IpAddress: s.IPAddress,
				Port:      s.Port,
			},
		},
	}, nil
}

func getLocalIPAddress() (net.IP, error) {
	conn, err := net.Dial("udp", "1.2.3.4:53")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
