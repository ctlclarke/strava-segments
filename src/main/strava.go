package main

import (
  "fmt"
  "../api"
  r "../response_bodies"
  "strconv"
)
func main() {

  var activity r.Activity = api.GetActivity("1720473515")
  segment_efforts := activity.Segment_efforts

  for _, segment_effort := range segment_efforts {
  	fmt.Printf(segment_effort.Segment.Name + ": ")
  	segmentId := segment_effort.Segment.Id
  	leaderboard := api.GetSegmentLeaderboard(strconv.Itoa(segmentId))
  	for _, entry := range leaderboard.Entries {
        if entry.Athlete_name == "Charles C." {
            fmt.Printf(strconv.Itoa(entry.Rank))
        }
  	}
  	fmt.Printf(" out of ")
  	fmt.Printf(strconv.Itoa(leaderboard.Entry_count) + "\n")
  }
}
