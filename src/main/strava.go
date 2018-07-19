package main

import (
  "fmt"
  "../api"
)
func main() {

  // var stats = api.GetAthleteStats("2316892")
  // fmt.Println(stats)

  fmt.Println(api.GetMyActivities())
}
