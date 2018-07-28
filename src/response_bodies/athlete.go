package response_bodies

type Athlete struct {
    Id                      int `json:"id"`
    Username                string `json:"username"`
    Resource_state          int `json:"resource_state"`
    Firstname               string `json:"firstname"`
    Lastname                string `json:"lastname"`
    City                    string `json:"city"`
    State                   string `json:"state"`
    Country                 string `json:"country"`
    Sex                     string `json:"sex"`
    Premium                 bool `json:"premium"`
    Created_at              string `json:"created_at"`
    Updated_at              string `json:"updated_at"`
    Badge_type_id           int `json:"badge_type_id"`
    Profile_medium          string `json:"profile_medium"`
    Profile                 string `json:"proifle"`
    Friend                  bool `json:"friend"`
    Follower                bool `json:"follower"`
    Email                   string `json:"email"`
    Follower_count          int `json:"follower_count"`
    Friend_count            int `json:"friend_count"`
    Mutual_friend_count     int `json:"mutual_friend_count"`
    Athlete_type            int `json:"athlete_type"`
    Date_preference         string `json:"date_preference"`
    Measurement_preference  string `json:"measurement_preference"`
    Clubs                   []Club `json:"clubs"`
    Ftp                     int `json:"ftp"`
    Weight                  float32 `json:"weight"`
    Bikes                   []Bike `json:"bikes"`
    Shoes                   []Shoe `json:"shoes"`
}

type Bike struct {
  Id              string `json:"id"`
  Primary         bool `json:"primary"`
  Name            string `json:"name"`
  Resource_state  int `json:"resource_state"`
  Distance        float32 `json:"distance"`
}

type Shoe struct {
  Id              string `json:"id"`
  Primary         bool `json:"primary"`
  Name            string `json:"name"`
  Resource_state  int `json:"resource_state"`
  Distance        float32 `json:"distance"`
}
