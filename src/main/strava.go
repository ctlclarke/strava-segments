package main

import (
  "fmt"
  "../api"
  r "../response_bodies"
)
func main() {

  // api.GetAthleteStats("2316892")
  // fmt.Println(stats)
  // api.GetActivity("1599518530")
  // var athlete r.Athlete = api.GetThisAthlete()

  // fmt.Println(athlete.Firstname)

  // segment := api.GetSegment("1672847")

  // fmt.Println(segment.Effort_count)

  var activity r.Activity = api.GetActivity("1720473515")

  fmt.Println(activity)
  // fmt.Println(api.GetMyActivities())



  // get activity, take segment_efforts
  // find where pr_rank is not null
}
