package main

import (
  "net"
  "fmt"
  "bufio"
  "os"
)

func main() {
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")

  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Send text: ")
    text, _ := reader.ReadString('\n')
    fmt.Fprintf(conn, text + "\n")

    msg, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: " + msg)
  }
}