package main

import (
	"log"
	"net"

	"github.com/Javlopez/gprc_calculator/proto"
	"github.com/Javlopez/gprc_calculator/service"
	"google.golang.org/grpc"
)

func main() {
	port := ":4040"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("start server on port: %s", port)

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, service.NewAddServer())

	err = srv.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
