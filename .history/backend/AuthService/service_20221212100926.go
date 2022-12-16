package main

import (
	"context"

	"github.com/golang/protobuf/proto"
)

type authServer struct {}

func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*pro)