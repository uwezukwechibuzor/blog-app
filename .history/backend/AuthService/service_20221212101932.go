package main

import (

	"google.golang.org/grpc"
)

type authServer struct {}


func maim() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, authServer{})
}