package api
import (
	r "../response_bodies"
  "encoding/json"
)
// func GetAthleteStats(athleteId string) r.Athlete {
//   url := "https://www.strava.com/api/v3/athletes/" + athleteId + "/stats"
//   var athlete r.Athlete
//   err := json.Unmarshal(get(url), &athlete)
//     if err != nil {
//         panic(err)
//     }
//   return athlete
// }

func GetThisAthlete() r.Athlete {
  url := "https://www.strava.com/api/v3/athlete/"
  var athlete r.Athlete
  err := json.Unmarshal(get(url), &athlete)
    if err != nil {
        panic(err)
    }
  return athlete
}

// // func GetSegment(segmentId string) response_bodies.Club {
// //   url := "https://www.strava.com/api/v3/segments/" + segmentId
// //   return get(url)
// // }

// func GetActivity(activityId string) r.Athlete {
//   url := "https://www.strava.com/api/v3/activities/" + activityId
//   return get(url) // TODO: qq: need to work on the 'get' method, could do with it returning a generic
// }

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
