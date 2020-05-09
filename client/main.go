package main

import (
	"context"
	"flag"
	"log"

	"github.com/Javlopez/gprc_calculator/proto"
	"google.golang.org/grpc"
)

func main() {
	operation := flag.String("operation", "sum", "sum")
	flag.Parse()

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	client := proto.NewAddServiceClient(conn)
	a := int64(5)
	b := int64(18)
	req := &proto.Request{A: a, B: b}

	var res *proto.Response

	if *operation == "-" {
		res, err = client.Substract(context.Background(), req)
	} else {
		res, err = client.Add(context.Background(), req)
	}
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Unnary request finished, response from the server:SUM %d\n", res.Result)

}
