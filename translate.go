package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
)

type gtresponse struct {
  Data struct {
    Detections [][]struct {
      Language string `json:"language"`
    } `json:"detections"`
  } `json:"data"`
}

func Detect_language(q string) string {
  var Url *url.URL
  Url, err := url.Parse("https://www.googleapis.com")
  if err != nil {
    fmt.Println(err)
  }

  Url.Path += "/language/translate/v2/detect"
  parameters := url.Values{}
  parameters.Add("q", q)
  parameters.Add("key", Google_key())
  Url.RawQuery = parameters.Encode()

  resp, err := http.Get(Url.String())
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  var data gtresponse
  json.Unmarshal(body, &data)

  lang := data.Data.Detections[0][0].Language
  return lang
}
