package main

import (
	"log"

	minio "github.com/minio/minio-go"
)

var client *minio.Client

func connect() {
	var err error

	accessKey := config.DOAccessKey
	secretKey := config.DOSecretKey
	if accessKey == "" || secretKey == "" {
		log.Fatal("Must set DO_ACCESS_KEY and DO_SECRET_KEY environment variables")
	}
	ssl := true

	client, err = minio.New("nyc3.digitaloceanspaces.com", accessKey, secretKey, ssl)
	if err != nil {
		log.Fatal(err)
	}
}

func checkKey(key, bucket string) bool {
	return true
}
