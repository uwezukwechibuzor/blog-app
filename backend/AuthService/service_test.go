package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/uwezukwechibuzor/blog-application/proto"
)

func Test_authServer_Login(t *testing.T) {
	type args struct {
		in0 context.Context
		in  *proto.LoginRequest
	}
	tests := []struct {
		name    string
		a       authServer
		args    args
		want    *proto.AuthResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := authServer{}
			got, err := a.Login(tt.args.in0, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("authServer.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authServer.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authServer_UsernameUsed(t *testing.T) {
	type args struct {
		in0 context.Context
		in  *proto.UsernameUsedRequest
	}
	tests := []struct {
		name    string
		a       authServer
		args    args
		want    *proto.UsedResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := authServer{}
			got, err := a.UsernameUsed(tt.args.in0, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("authServer.UsernameUsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authServer.UsernameUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authServer_EmailUsed(t *testing.T) {
	type args struct {
		in0 context.Context
		in  *proto.EmailUsedRequest
	}
	tests := []struct {
		name    string
		a       authServer
		args    args
		want    *proto.UsedResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := authServer{}
			got, err := a.EmailUsed(tt.args.in0, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("authServer.EmailUsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authServer.EmailUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}
