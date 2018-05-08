package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
)

func main() {
  resp, err := http.Get("https://api.ipify.org")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
}
