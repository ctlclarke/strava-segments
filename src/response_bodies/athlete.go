package response_bodies

type Athlete struct {
    id                string
    username              string
    resource_state        int
    firstname           string
    lastname            string
    city          string
    state           string
    country         string
    sex           string
    premium         bool
    created_at        string
    updated_at        string
    badge_type_id     int
    profile_medium      string
    profile         string
    friend          bool
    follower        bool
    email           string
    follower_count      int
    friend_count      int
    mutual_friend_count   int
    athlete_type      int
    date_preference     string
    measurement_preference  string
    clubs         []Club
    ftp           int
    weight          int
    bikes         []Bike
    shoes         []Shoe
}

type Bike struct {
  id            string
  primary         bool
  name          string
  resource_state      int
  distance        int
}

type Shoe struct {
  id            string
  primary         bool
  name          string
  resource_state      int
  distance        int
}
