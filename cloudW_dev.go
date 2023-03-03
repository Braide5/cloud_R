package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function for the root path
    rootHandler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
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