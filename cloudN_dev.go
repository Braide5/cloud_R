package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    // Get the command line arguments
    args := os.Args[1:]

    // Check if any arguments were provided
    if len(args) == 0 {
        fmt.Println("Please provide a command")
        return
    }

    // Get the command
    command := args[0]

    // Switch based on the command
    switch command {
    case "ping":
        // Code to ping a host
        if len(args) < 2 {
            fmt.Println("Please provide a host to ping")
            return
        }
        host := args[1]
        conn, err := net.Dial("ip4:icmp", host)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer conn.Close()
        fmt.Printf("Pinging %s...\n", host)
    case "traceroute":
        // Code to trace the route to a host
        if len(args) < 2 {
            fmt.Println("Please provide a host to trace")
            return
        }
        host := args[1]
        fmt.Printf("Tracing route to %s...\n", host)
        addrs, err := net.LookupIP(host)
        if err != nil {
            fmt.Println(err)
            return
        }
        for _, addr := range addrs {
            fmt.Println(addr)
        }
    default:
        // Invalid command
        fmt.Printf("Invalid command: %s\n", command)
    }
}
//Go program that can be used to develop network tools
//understand and modify
//please provide a command