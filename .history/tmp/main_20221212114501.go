package main

import (
	"context"

	"github.com/uwezukwechibuzor/blog-application/global"
)

func main() {
	global.DB.Collection("user").InsertOne(context.Background())
}