package api

import (
  "log"
  "net/http"
  "io/ioutil"
  "os"
  "encoding/json"
   "../response_bodies"
)

type Foo struct {
    id string
}

func get(url string) response_bodies.Athlete{

  var token string = os.Getenv("STRAVA_TOKEN")
  var athlete response_bodies.Athlete

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
  }
  client := &http.Client{}
  req.Header.Add("Authorization", "Bearer " + token)
  req.Header.Set("Content-Type", "application/json")
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal("Do: ", err)
  }

  defer resp.Body.Close()
  if resp.StatusCode == http.StatusOK {
    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    err := json.Unmarshal(bodyBytes, &athlete)
    if err != nil {
        panic(err)
    }
  }
  return athlete
}
