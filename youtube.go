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
    ContentDetails struct {
      RegionRestriction struct {
        Allowed []string `json:"allowed"`
        Blocked []string `json:"blocked"`
      } `json:"regionRestriction"`
    } `json:"contentDetails"`
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
	NextPageToken string `json:"nextPageToken"`
}

func Retrieve_all_videos(token string) {
  regions := []string{"en-US", "es-ES", "fr-FR", "sv-SE"}
  Videos = map[string][]string{}

  for _, r := range regions {
    language := strings.Split(r, "-")[0]
    region := strings.Split(r, "-")[1]

    Videos[language] = []string{}
    Retrieve_videos(region, language, "")
  }

  fmt.Println("Done retrieving videos")
}

func Retrieve_videos(region, lang, token string) {
  s := []string{}
  s = append(s, "https://www.googleapis.com/youtube/v3/videos")
  s = append(s, "?part=id,contentDetails,snippet")
  s = append(s, "&chart=mostPopular")
  s = append(s, "&videoCategoryId=10")
  s = append(s, "&regionCode=")
  s = append(s, region)
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

    allowed := video.ContentDetails.RegionRestriction.Allowed
    blocked := video.ContentDetails.RegionRestriction.Blocked

    if len(allowed) != 0 {
      if !stringInSlice("US", allowed) {
        break
      }
    }

    if len(blocked) != 0 {
      if stringInSlice("US", blocked) {
        break
      }
    }

    detected_lang := Detect_language(title)
    if detected_lang == lang {
      Videos[lang] = append(Videos[lang], id)
    }
  }

  nextToken := data.NextPageToken
  if nextToken != "" {
    Retrieve_videos(region, lang, nextToken)
  }
}
