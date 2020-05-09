package service

import (
	"context"

	"github.com/Javlopez/gprc_calculator/proto"
)

type AddServer struct {
}

func NewAddServer() *AddServer {
	return &AddServer{}
}

func (srv *AddServer) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (srv *AddServer) Substract(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a - b

	return &proto.Response{Result: result}, nil
}
