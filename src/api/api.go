package api

import (
  "log"
  "net/http"
  "io/ioutil"
  "os"
  "fmt"
  "encoding/json"
   "../response_bodies"
)

type Foo struct {
    id string
}

func get(url string) string{

  var token string = os.Getenv("STRAVA_TOKEN")

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    // return ""
  }
  client := &http.Client{}
  req.Header.Add("Authorization", "Bearer " + token)
  req.Header.Set("Content-Type", "application/json")
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal("Do: ", err)
    // return ""
  }

  defer resp.Body.Close()
  var bodyString string
  if resp.StatusCode == http.StatusOK {
    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    bodyString = string(bodyBytes)
  
  var dat map[string]interface{}

  var athlete response_bodies.Athlete

  err := json.Unmarshal(bodyBytes, &athlete)
    if err != nil {
      fmt.Println("whoops 1")
        panic(err)
    }
    fmt.Println(athlete.Bikes)

  if err := json.Unmarshal(bodyBytes, &dat); err != nil {
    fmt.Println("whoops 2")
        panic(err)
  }

  }
  return bodyString
}
