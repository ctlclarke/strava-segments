package api

import (
  "log"
  "net/http"
  "io/ioutil"
)

func get(url string) string {
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    return ""
  }
  client := &http.Client{}
  req.Header.Add("Authorization", "Bearer 6651096a94211f2e4a73d4a36db6a35066382adc")
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal("Do: ", err)
    return ""
  }

  defer resp.Body.Close()
  var bodyString string

  if resp.StatusCode == http.StatusOK {
    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    bodyString = string(bodyBytes)
  }

  return bodyString
}
