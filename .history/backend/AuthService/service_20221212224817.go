package main

import (
	"context"
	"errors"
	"log"
	"net"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

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
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return &proto.AuthResponse{}, errors.New("Wrong login Credentials Provided")
	}
	return &proto.AuthResponse{Token: user.GetToken()}, nil
}


func (authServer) Signup(_ context.Context, in *proto.SignupRequest) (*proto.AuthResponse, error) {
	username, email, password := in.GetUsername(), in.GetEmail(), in.GetPassword()
	match, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	if len(username) < 4 || len(username) >  {
		return &proto.AuthResponse{}, errors.New("validation failed")
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
