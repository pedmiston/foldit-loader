package main

import (
	"flag"
	"fmt"
)

func main() {
	var keys chan string

	bucket := flag.String("bucket", "foldit", "Name of bucket holding data.")
	key := flag.String("key", "", "Key for specific data file to load. Optional.")

	if *key != "" {
		keys = make(chan string)
		keys <- *key
	} else {
		keys = loadKeys(*bucket)
	}

	for key := range keys {
		fmt.Println(key)
	}
}
