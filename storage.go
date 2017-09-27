package main

import (
	"log"
	"os"

	minio "github.com/minio/minio-go"
)

var client *minio.Client

func connect() {
	var err error

	accessKey := os.Getenv("DO_ACCESS_KEY")
	secretKey := os.Getenv("DO_SECRET_KEY")
	if accessKey == "" || secretKey == "" {
		log.Fatal("Must set DO_ACCESS_KEY and DO_SECRET_KEY environment variables")
	}
	ssl := true

	client, err = minio.New("nyc3.digitaloceanspaces.com", accessKey, secretKey, ssl)
	if err != nil {
		log.Fatal(err)
	}
}
