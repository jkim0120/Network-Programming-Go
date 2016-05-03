package main

import (
  "net"
  "fmt"
  "bufio"
  "strings"
)

func main() {
  fmt.Println("Starting server...")

  ln, _ := net.Listen("tcp", ":8081")
  conn, _ := ln.Accept()

  for {
    msg, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message received:", string(msg))

    newmsg := strings.ToUpper(msg)
    conn.Write([]byte(newmsg + "\n"))
  }
}