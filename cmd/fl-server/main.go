package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
    //Here we are using the net.Conn because it is just a go wrapper around OS - Lwevel Socket
    //For example , using the net.Conn, we could speak using conn.Write() and listen using conn.Read()
    //Establishing a live connection over Wifi between Clients RAM and Server's RAM  
    buffer := make([]byte,1024)

    n,err := conn.Read(buffer)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Converting the buffer into byte slices
    buffer_byte := string(buffer[:n])
    fmt.Printf(buffer_byte)

    
    //Closing the connection as its directly connected with the OS resurces , just to prevent the Memory leaks
    conn.Close()
}

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

        go handleConnection(conn)
        // the handleConnection handles the connection , so the close Connection is not required here anymore
    }
}