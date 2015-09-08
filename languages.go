package main

import (
  "encoding/json"
  "io/ioutil"
)

type Language struct {
  Movies string
}

func Language_file(lang string) Language {
  data, _ := ioutil.ReadFile("languages/" + lang + ".json")
  var language Language

  json.Unmarshal(data, &language)
  return language
}

func Movies_url(lang string) string {
  language := Language_file(lang)
  return language.Movies
}
