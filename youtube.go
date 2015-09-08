package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
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
  var Url *url.URL
  Url, err := url.Parse("https://www.googleapis.com")
  if err != nil {
    fmt.Println(err)
  }

  Url.Path += "/youtube/v3/videos"
  parameters := url.Values{}
  parameters.Add("part", "id,contentDetails,snippet")
  parameters.Add("chart", "mostPopular")
  parameters.Add("videoCategoryId", "10")
  parameters.Add("regionCode", region)
  parameters.Add("pageToken", token)
  parameters.Add("key", Google_key())
  Url.RawQuery = parameters.Encode()

  resp, err := http.Get(Url.String())
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
