package main

import (
  "fmt"
  "math/rand"
  "net/http"
  "time"
)

var Videos map[string][]string

func main() {
  Retrieve_all_videos(Google_key())

  ticker := time.NewTicker(1 * time.Hour)
  quit := make(chan struct{})
  go func() {
    for {
      select {
      case <- ticker.C:
        Retrieve_all_videos(Google_key())
      case <- quit:
        ticker.Stop()
        return
      }
    }
  }()

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    http.ServeFile(w, r, path)
  })

  http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
    lang := r.URL.Query().Get("lang")

    fmt.Fprintf(w, "<html><meta http-equiv=\"refresh\" content=\"0;URL=%s\"></html>", Movies_url(lang))
  })

  http.HandleFunc("/music", func(w http.ResponseWriter, r *http.Request) {
    lang := r.URL.Query().Get("lang")

    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    index := r1.Intn(len(Videos[lang]))
    redirect := Videos[lang][index]
    fmt.Fprintf(w, "<html><meta http-equiv=\"refresh\" content=\"0;URL=https://youtu.be/%s\"></html>", redirect)
  })

  http.ListenAndServe(":8080", nil)
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
