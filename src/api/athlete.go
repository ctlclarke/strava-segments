package api
import (
	r "../response_bodies"
  "encoding/json"
  "fmt"
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
  var x = get(url)
  fmt.Println(x)
  err := json.Unmarshal(get(url), &segment)
  if err != nil {
    panic(err)
  }
  return segment
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

// func GetMyActivities() response_bodies.Club {
//   url := "https://www.strava.com/api/v3/athlete/activities"
//   return get(url)
// }

// func GetClub(clubId string, target interface{}){
  // url := "https://www.strava.com/api/v3/clubs/" + clubId
  // get(url)
// 
  // return
// }

// TODO: qq: get all these to return the objects, not the strings.
// TODO: qq: or re-objectify them in strava.go?
