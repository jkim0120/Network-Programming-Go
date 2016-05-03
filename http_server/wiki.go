package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

type Page struct {
  Title string
  Body []byte
}

type wikiHandler struct {}

func (*wikiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if handler, ok := mux[r.URL.String()]; ok {
    handler(w, r)
    return
  }

  viewHandler(w, r)
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
  // 0600 -> read write permission
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func main() {
  server := http.Server {
    Addr: ":8080",
    Handler: &wikiHandler{},
  }

  mux := make(map[string]func(http.ResponseWriter, *http.Request))
  mux["/view/"] = viewHandler

  server.ListenAndServe()
}