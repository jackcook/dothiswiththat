package main

import (
  "html/template"
  "io/ioutil"
  "net/http"
)

type Page struct {
    Title string
    Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
  // p := loadPage(title)
  p := &Page{Title: "test"}
  t, _ := template.ParseFiles("index.html")
  t.Execute(w, p)
}

func loadPage(title string) *Page {
    filename := title + ".txt"
    body, _ := ioutil.ReadFile(filename)
    return &Page{Title: title, Body: body}
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
