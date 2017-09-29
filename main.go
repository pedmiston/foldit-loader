package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db     *gorm.DB
	config Config
)

func main() {
	var keys chan string

	bucket := flag.String("bucket", "foldit", "Name of bucket holding data.")
	key := flag.String("key", "", "Key for specific data file to load. Optional.")
	creds := flag.String("creds", "", "File containing secrets")

	config = readCreds(*creds)

	if *key != "" {
		ok := checkKey(*key, *bucket)
		if !ok {
			fmt.Fprintf(os.Stderr, "key '%v' not found in bucket '%v'", *key, *bucket)
			os.Exit(1)
		}
		keys = make(chan string)
		keys <- *key
	} else {
		keys = loadKeys(*bucket)
	}

	for key := range keys {
		fmt.Println(key)
	}

	//dsn := fmt.Sprintf("%s:%s@%s:%s/%s", username, password, host, port, dbname)
	//db, err = gorm.Open("username:password@protocol(address)/dbname?param=value")
	//defer db.Close()
}
