package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "math/rand"
  "net/http"
  "strings"
  "time"
)

type Response struct {
  Tracks Tracks
}

type Tracks struct {
  Href string
  Items []Track
  Limit int
  Next string
  Offset int
  Total int
}

type Track struct {
  Available_markets []string
  External_urls map[string]string
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    http.ServeFile(w, r, path)
  })

  http.HandleFunc("/music", func(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    market := r.URL.Query().Get("market")

    s := []string{}
    s = append(s, "https://api.spotify.com/v1/search?type=track&q=")
    s = append(s, q)
    s = append(s, "&market=")
    s = append(s, market)

    url := strings.Join(s, "")

    resp, err := http.Get(url)
    if err != nil {
      fmt.Println(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    var redirect_urls []string
    var data Response
    json.Unmarshal(body, &data)
    for _, track := range data.Tracks.Items {
      if stringInSlice("US", track.Available_markets) {
        redirect_url := track.External_urls["spotify"]
        redirect_urls = append(redirect_urls, redirect_url)
      }
    }

    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    index := r1.Intn(len(redirect_urls))
    redirect := redirect_urls[index]
    fmt.Fprintf(w, "<html><meta http-equiv=\"refresh\" content=\"0;URL=%s\"></html>", redirect)
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
