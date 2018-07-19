package api

func GetAthleteStats(athleteId string) string {
  url := "https://www.strava.com/api/v3/athletes/" + athleteId + "/stats"
  return get(url)
}

func GetSegment(segmentId string) string {
  url := "https://www.strava.com/api/v3/segments/" + segmentId
  return get(url)
}

func GetActivity(activityId string) string {
  url := "https://www.strava.com/api/v3/activities/" + activityId
  return get(url)
}

func GetMyActivities() string {
  url := "https://www.strava.com/api/v3/athlete/activities"
  return get(url)
}
