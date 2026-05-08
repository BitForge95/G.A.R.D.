package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")

    if err != nil {
        fmt.Println("THere is some error while conneting to the PORT")
        os.Exit(1)
    }

    fmt.Println("Server is listening on port 8080")

    for {
        conn, err := ln.Accept()

        if err != nil {
            fmt.Println("There's something wrong with connection")
            continue
        }

        fmt.Println("New Clinet connected")
        conn.Close()
    }
}