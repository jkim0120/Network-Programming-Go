package main

import (
  "io"
  "net/http"
)

type myHandler struct {}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if handler, ok := mux[r.URL.String()]; ok {
    handler(w, r)
    return
  }

  io.WriteString(w, "My server: " + r.URL.String())
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func handler(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello World!")
}

func main() {
  server := http.Server {
    Addr: ":8080",
    Handler: &myHandler{},
  }

  mux := make(map[string]func(http.ResponseWriter, *http.Request))
  mux["/"] = handler

  server.ListenAndServe()
}