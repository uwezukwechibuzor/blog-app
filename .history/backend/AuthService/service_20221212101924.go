package main

import (
	"context"

	"github.com/uwezukwechibuzor/blog-application/proto"
	"google.golang.org/grpc"
)

type authServer struct {}


func maim() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, authServer{})
}