package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestListKeysInBucket(t *testing.T) {
	setEnv("secrets.json")
	connect()
}

func setEnv(secrets string) {
	f, err := ioutil.ReadFile(secrets)
	if err != nil {
		log.Fatal(err)
	}
	a := make(map[string]string)
	json.Unmarshal(f, &a)
	fmt.Println(a)
	os.Setenv("DO_ACCESS_KEY", a["accessKey"])
	os.Setenv("DO_SECRET_KEY", a["secretKey"])
}
