package api

import (
  "log"
  "net/http"
  "io/ioutil"
  "os"
  "fmt"
  "encoding/json"
)

func get(url string) string {

  var token string = os.Getenv("STRAVA_TOKEN")

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    return ""
  }
  client := &http.Client{}
  req.Header.Add("Authorization", "Bearer " + token)
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
  
  var dat map[string]interface{}

  // fmt.Println(dat)

  // num := dat["id"]

  // fmt.Println(num)
  if err := json.Unmarshal(bodyBytes, &dat); err != nil {
        panic(err)
  }

  // var seggers map[string]interface{}

  // if err := json.Unmarshal(dat["segment_efforts"], &seggers); err != nil {
  //       panic(err)
  // } 

  // fmt.Println(seggers)
  // }
  fmt.Println(dat["segment_efforts"])
  }
  return bodyString
}
