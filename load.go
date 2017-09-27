package main

import "fmt"

// Load object keys in the given bucket into a channel.
func loadKeys(bucket string) chan string {
	keys := make(chan string)

	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	var prefix string
	isRecursive := true
	objectCh := client.ListObjects(bucket, prefix, isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return keys
		}
		keys <- object.Key
	}

	return keys
}
