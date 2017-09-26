package main

func main() {
	bucket := flag.String("bucket", "foldit", "name of bucket holding data")
	key := flag.String("key", "", "key for data file to load")

	keys := make(chan string)

	if key != "" {
		keys <- *key
	} else {
		loadKeys(*bucket, keys)
	}

	for key := range keys {
		err := Upload(*bucket, key)
		if err != nil {
			log.Fatal(err)
		}
	}
}
