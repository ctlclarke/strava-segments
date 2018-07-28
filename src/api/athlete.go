package api
import (
	// "../response_bodies"
)
func GetAthleteStats(athleteId string) string {
  url := "https://www.strava.com/api/v3/athlete/"// + athleteId + "/stats" TODO: qq: this should be athlete__S__
  return get(url)
}

// func GetSegment(segmentId string) response_bodies.Club {
//   url := "https://www.strava.com/api/v3/segments/" + segmentId
//   return get(url)
// }

func GetActivity(activityId string) string {
  url := "https://www.strava.com/api/v3/activities/" + activityId
  return get(url)
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
