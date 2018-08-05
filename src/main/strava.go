package main

import (
  "fmt"
  "../api"
  r "../response_bodies"
  "strconv"
  "sort"
)
func main() {

  // listSegmentsByPercentile("1642908894", 30)

  fmt.Println("Max speed in last 30 activities: " + strconv.FormatFloat(getMaxSpeed() * 3.6, 'g', 6, 64) + "km/h")
}

func getMaxSpeed() float64{
  activities := api.GetMyActivities() // This only gets the 30 most recent activities
  var max_speed float64 = 0

  for _, activity := range activities {
    if (activity.Max_speed > max_speed) {
      max_speed = activity.Max_speed
    }
  }

  return max_speed
}

func listSegmentsByPercentile(activityId string, percentile float64) {
  segments := getSegmentRankingsForActivity(activityId, "Charles C.")

  for _, seg := range segments {
    if seg.Percentile < 30 {
      fmt.Println(seg.SegmentName + " : " + strconv.FormatFloat(seg.Percentile, 'g', 6, 64))
      } else {
        break
      }
  }
}

func getSegmentRankingsForActivity(activityId string, athleteName string) []SegmentRankings{
    var activity r.Activity = api.GetActivity(activityId)
    var leaderboardPosition int = -1
    var percentile float64 = -1

    segment_efforts := activity.Segment_efforts

    segmentRankings := make([]SegmentRankings, len(activity.Segment_efforts))

    for index, segment_effort := range segment_efforts {
      segmentRankings[index].SegmentId = strconv.Itoa(segment_effort.Id)
      segmentRankings[index].SegmentName = segment_effort.Segment.Name

      segmentId := segment_effort.Segment.Id
      leaderboard := api.GetSegmentLeaderboard(strconv.Itoa(segmentId))

      for _, entry := range leaderboard.Entries {
        if entry.Athlete_name == athleteName {
            leaderboardPosition = entry.Rank
        }
      }
      percentile = 100.0 * float64(leaderboardPosition) / float64(leaderboard.Entry_count)

      segmentRankings[index].Position = leaderboardPosition
      segmentRankings[index].TotalEfforts = leaderboard.Entry_count
      segmentRankings[index].Percentile = percentile

    }

    sort.Slice(segmentRankings, func(i, j int) bool {
      return segmentRankings[i].Percentile < segmentRankings[j].Percentile
    })

    return segmentRankings
}

type SegmentRankings struct {
  SegmentId     string
  SegmentName   string
  Position      int
  TotalEfforts  int
  Percentile    float64
}