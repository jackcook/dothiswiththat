package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
)

type ytresponse struct {
	Items []struct {
		ID string `json:"id"`
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
	NextPageToken string `json:"nextPageToken"`
}

func Retrieve_videos(token string) {
  s := []string{}
  s = append(s, "https://www.googleapis.com/youtube/v3/videos")
  s = append(s, "?part=snippet")
  s = append(s, "&chart=mostPopular")
  s = append(s, "&videoCategoryId=10")
  s = append(s, "&regionCode=se")
  s = append(s, "&pageToken=")
  s = append(s, token)
  s = append(s, "&key=")
  s = append(s, Google_key())

  url := strings.Join(s, "")

  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  var data ytresponse
  json.Unmarshal(body, &data)

  for _, video := range data.Items {
    id := video.ID
    title := video.Snippet.Title
    lang := Detect_language(title)
    if lang == "sv" {
      Videos = append(Videos, id)
    }
  }

  nextToken := data.NextPageToken
  if nextToken != "" {
    Retrieve_videos(nextToken)
  }
}
