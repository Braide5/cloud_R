package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
)

func main() {
	// Set up a client to access Google Cloud Storage
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Set up an HTTP handler to serve a file from Google Cloud Storage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bucketName := os.Getenv("BUCKET_NAME")
		objectName := os.Getenv("OBJECT_NAME")

		reader, err := client.Bucket(bucketName).Object(objectName).NewReader(r.Context())
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}
		defer reader.Close()

		_, err = io.Copy(w, reader)
		if err != nil {
			log.Fatalf("Failed to copy file to response: %v", err)
		}
	})

	// Start the server and listen for incoming requests
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
//a go program that demonstrates how to build a simple cloud application using Google Cloud Platform
//understand and modify