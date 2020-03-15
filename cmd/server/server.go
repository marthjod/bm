package main

import (
	"fmt"

	"github.com/marthjod/bm/rpc/proto"
	"github.com/marthjod/bm/rpc/server"
	"google.golang.org/grpc"

	"log"
	"net"
)

const (
	port = 50051
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", srv)
	s := grpc.NewServer()
	proto.RegisterVersionerServer(s, srv)
	log.Printf("listening on :%d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
