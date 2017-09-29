package main

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DOAccessKey   string `yaml:"do_access_key"`
	DOSecretKey   string `yaml:"do_secret_key"`
	MySQLUsername string `yaml:"mysql_username"`
	MySQLPassword string `yaml:"mysql_password"`
}

func readCreds(creds string) Config {
	f, err := os.Open(creds)
	if err != nil {
		log.Fatal(err)
	}

	var b []byte
	_, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	c := &Config{}
	err = yaml.Unmarshal(b, c)
	if err != nil {
		log.Fatal(err)
	}

	return *c
}
