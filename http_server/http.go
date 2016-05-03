package main

import (
  "io"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello World!")
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", handler)
  http.ListenAndServe(":8080", mux)
}