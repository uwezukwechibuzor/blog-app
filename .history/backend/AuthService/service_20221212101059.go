package main

import (
	"context"

	"github.com/uwezukwechibuzor/blog-application/proto"
)

type authServer struct {}

func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
  return
}

func maim() {
	server := grpc>
}