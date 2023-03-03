package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()

	// Set up a client to access Google Cloud Storage
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		return
	}

	// List all the buckets in your Google Cloud Storage project
	buckets, err := client.Buckets(ctx, "your-project-id")
	if err != nil {
		fmt.Printf("Failed to list buckets: %v\n", err)
		return
	}

	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}
}
//just a golang program related to cloud
//understand every line and modify itS
//1st.go:7:2: no required module provides package cloud.google.com/go/storage: go.mod file not found in current directory or any parent directory; see 'go help modules'