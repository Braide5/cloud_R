package main

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
)

func main() {
    // Define a handler function for the root path
    rootHandler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
        // Set the seed value for the random number generator
    rand.Seed(time.Now().UnixNano())

    // Generate and print 10 random integers between 0 and 99
    for i := 0; i < 10; i++ {
        fmt.Println(rand.Intn(100))
    }
}

    // Define a handler function for the "/about" path
    aboutHandler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "About me!")
    }
    infoHandler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Information about me!")
    }
    contactHandler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "contact me!")
    }
    
    // Register the handlers
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/info", infoHandler)
    http.HandleFunc("/contact", contactHandler)

    // Start the server
    fmt.Println("Starting server on port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }
}
//Go program that can be used for web development
//understand and modify
//please provide a command
//this modified one givees output on the terminal