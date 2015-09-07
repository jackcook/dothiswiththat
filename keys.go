package main

import (
  "encoding/json"
  "io/ioutil"
)

type Keys struct {
  Googleapis string
}

func Google_key() string {
  data, _ := ioutil.ReadFile("keys.json")
  var keys Keys

  json.Unmarshal(data, &keys)
  return keys.Googleapis
}
