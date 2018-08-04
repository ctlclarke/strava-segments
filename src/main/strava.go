package main

import (
  "fmt"
  "../api"
  r "../response_bodies"
  "strconv"
)
func main() {

  var activity r.Activity = api.GetActivity("1720473837")
  segment_efforts := activity.Segment_efforts
  var entryRank int
  var percentile float64

  for _, segment_effort := range segment_efforts {
  	fmt.Printf(segment_effort.Segment.Name + ": ")
  	segmentId := segment_effort.Segment.Id
  	leaderboard := api.GetSegmentLeaderboard(strconv.Itoa(segmentId))
  	for _, entry := range leaderboard.Entries {
        if entry.Athlete_name == "Charles C." {
            entryRank = entry.Rank
        }
  	}
    percentile = 100.0 * float64(entryRank) / float64(leaderboard.Entry_count)
    fmt.Printf(strconv.Itoa(entryRank))
    fmt.Printf(" out of ")

    if percentile < 20 {
      fmt.Printf(strconv.Itoa(leaderboard.Entry_count) + ":  \t\t\t\t  ")
      fmt.Printf(strconv.FormatFloat(percentile, 'g', 6, 64) + "\n")
    } else {
            fmt.Printf(strconv.Itoa(leaderboard.Entry_count) + "\n")
    }

  }
}
