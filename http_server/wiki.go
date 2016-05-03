package main

import (
  "strings"
  "html/template"
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

  action := strings.Split(r.URL.Path, "/")
  switch action[1] {
    case "view":
      viewHandler(w, r)
    case "edit":
      editHandler(w, r)
    case "save":
      saveHandler(w, r)
  }
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  t, err := template.ParseFiles(tmpl + ".html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = t.Execute(w, p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w, r, "/edit/" + title, http.StatusFound)
    return
  }
  renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/save/"):]
  body := r.FormValue("body")
  p := &Page{
    Title: title,
    Body: []byte(body),
  }
  err := p.save()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  http.Redirect(w, r, "/view/" + title, http.StatusFound)
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
  mux["/edit/"] = editHandler
  mux["/save/"] = saveHandler

  server.ListenAndServe()
}