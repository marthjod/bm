package main

import (
	"github.com/marthjod/bm/rpc/proto"
	"github.com/marthjod/bm/rpc/server"
	"google.golang.org/grpc"

	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterVersionerServer(s, &server.Server{})
	log.Printf("listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
