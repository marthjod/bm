package server

import (
	"context"

	"github.com/marthjod/bm/rpc/proto"
)

type Server struct {
	proto.VersionerServer
}

// Version returns the server's version.
func (s *Server) Version(ctx context.Context, req *proto.VersionRequest) (*proto.VersionResponse, error) {
	return &proto.VersionResponse{
		Version: &proto.Version{},
	}, nil
}
