package main

import (
  "fmt"
  "math/rand"
  "net/http"
  "time"
)

var Videos []string

func main() {
  Retrieve_videos("")
  for _, id := range Videos {
    fmt.Println(id)
  }

  ticker := time.NewTicker(1 * time.Hour)
  quit := make(chan struct{})
  go func() {
    for {
      select {
      case <- ticker.C:
        Retrieve_videos("")
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

  http.HandleFunc("/music", func(w http.ResponseWriter, r *http.Request) {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    index := r1.Intn(len(Videos))
    redirect := Videos[index]
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
