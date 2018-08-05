package api
import (
	r "../response_bodies"
  "encoding/json"
)
func GetAthleteStats(athleteId string) r.Athlete {
  url := "https://www.strava.com/api/v3/athletes/" + athleteId + "/stats"
  var athlete r.Athlete
  err := json.Unmarshal(get(url), &athlete)
  if err != nil {
      panic(err)
  }
  return athlete
}

func GetThisAthlete() r.Athlete {
  url := "https://www.strava.com/api/v3/athlete/"
  var athlete r.Athlete
  err := json.Unmarshal(get(url), &athlete)
  if err != nil {
      panic(err)
  }
  return athlete
}

func GetSegment(segmentId string) r.Segment { // TODO: qq: untested
  url := "https://www.strava.com/api/v3/segments/" + segmentId
  var segment r.Segment
  err := json.Unmarshal(get(url), &segment)
  if err != nil {
    panic(err)
  }
  return segment
}

func GetSegmentLeaderboard(segmentId string) r.Leaderboard {
  url := "https://www.strava.com/api/v3/segments/" + segmentId + "/leaderboard"
  var leaderboard r.Leaderboard
  err := json.Unmarshal(get(url), &leaderboard)
  if err != nil {
    panic(err)
  }
  return leaderboard
}

func GetActivity(activityId string) r.Activity {
  url := "https://www.strava.com/api/v3/activities/" + activityId
  var activity r.Activity
  err := json.Unmarshal(get(url), &activity)
  if err != nil {
    panic(err)
  }
  return activity
}

func GetMyActivities() []r.Activity {
  url := "https://www.strava.com/api/v3/athlete/activities"
  var activities []r.Activity
  err := json.Unmarshal(get(url), &activities)
  if err != nil {
    panic(err)
  }
  return activities
}

// func GetClub(clubId string, target interface{}){
  // url := "https://www.strava.com/api/v3/clubs/" + clubId
  // get(url)
// 
  // return
// }
