// package main

import "fmt"
// import "rsc.io/quote"
import "net"

const port = ":8080"

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Printf("Connection established from %s\n", conn.RemoteAddr().String())

    buffer := make([]byte, 10024)
    n,err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading from connection:", err)
        return
    }
    message := string(buffer[:n])
    fmt.Printf("Received message: %s\n", message)

    response := "Hello, World!"
    _, err = conn.Write([]byte(response))
    if err != nil {
        fmt.Println("Error writing to connection:", err)
        return
    }
    conn.Close()
}
func startServer(cnt int) {
    ln, err := net.Listen("tcp", port)
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    fmt.Printf(`Server started on %s`, port)
    for{
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        cnt++;
        fmt.Println(cnt);
        go handleConnection(conn);
    }
}

func main() {
    startServer(0)
}