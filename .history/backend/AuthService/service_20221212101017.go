package main

import (
	"context"

	"github.com/golang/protobuf/prot"
)

type authServer struct {}

func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {

}