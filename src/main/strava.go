package main

import (
  "fmt"
  "../api"
  // "encoding/json"
  r "../response_bodies"
)
func main() {

  // api.GetAthleteStats("2316892")
  // fmt.Println(stats)
  // api.GetActivity("1599518530")
  var athlete r.Athlete = api.GetThisAthlete()

  fmt.Println(athlete.Firstname)
  // fmt.Println(api.GetMyActivities())
}
