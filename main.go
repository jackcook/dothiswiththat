package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    http.ServeFile(w, r, path)
  })

  http.ListenAndServe(":8080", nil)
}
