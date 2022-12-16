package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	"golang.org/"

	"github.com/uwezukwechibuzor/blog-application/global"
	"github.com/uwezukwechibuzor/blog-application/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type authServer struct{}

func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	login, password := in.GetLogin(), in.GetPassword()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var user global.User
	global.DB.Collection("user").FindOne(ctx, []bson.M{bson.M{"username": login}, bson.M{"email": login}}).Decode(&user)
	if user == global.NilUser {
		return &proto.AuthResponse{}, errors.New("Wrong login Credentials Provided")
	}
   

}

func maim() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, proto.UnimplementedAuthServiceServer{})
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Error creating listener: ", err.Error())
	}
	server.Serve(listener)
}
