package main

import (
	"flag"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var keys chan string
	var n int

	bucket := flag.String("bucket", "foldit", "Name of bucket holding data.")
	key := flag.String("key", "", "Key for specific data file to load. Optional.")
	creds := flag.String("creds", "secrets.yml", "File containing secrets")
	flag.Parse()

	authenticate(*creds)
	connectStorage()

	if *key != "" {
		keys = loadKey(*key, *bucket)
		n = 1
	} else {
		keys, n = loadKeysFromBucket(*bucket)
	}

	for i := 0; i < n; i++ {
		k := <-keys
		fmt.Println(k)
	}

	//dsn := fmt.Sprintf("%s:%s@%s:%s/%s", username, password, host, port, dbname)
	//db, err = gorm.Open("username:password@protocol(address)/dbname?param=value")
	//defer db.Close()
}
