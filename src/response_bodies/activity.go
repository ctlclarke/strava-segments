package response_bodies

type Activity struct {
    Id                      int     `json:"id"`
    Resource_state          int     `json:"resource_state"`
    External_id             string  `json:"external_id"`
    Upload_id               int     `json:"upload_id"`
    Athlete                 Activity_Pair `json:"athlete"`
    Name                    string  `json:"string"`
    Distance                float32     `json:"distance"`
    Moving_time             int     `json:"moving_time"`
    Elapsed_time            int     `json:"elapsed_time"`
    Total_elevation_gain    float32     `json:"total_elevation_gain"`
    Type                    string  `json:"type"`
    Start_date              string  `json:"start_date"`
    Start_date_local        string  `json:"start_date_local"`
    Timezone                string  `json:"timezone"`
    Utc_offset              float32     `json:"utc_offset"`
    Start_latlng            []float32 `json:"start_latlng"`
    End_latlng              []float32 `json:"end_latlng"`
    Location_city           string  `json:"location_city"`
    Location_state          string  `json:"location_state"`
    Location_country        string  `json:"location_country"`
    Start_latitude          float64 `json:"start_latitude"`
    Start_longitude         float64 `json:"start_longitude"`
    Achievment_count        int     `json:"achievment_count"`
    Kudos_count             int     `json:"kudos_count"`
    Comment_count           int     `json:"comment_count"`
    Athlete_count           int     `json:"athlete_count"`
    Photo_count             int     `json:"photo_count"`
    Map                     Map     `json:"map"`
    Trainer                 bool    `json:"trainer"`
    Commute                 bool    `json:"commute"`
    Manual                  bool    `json:"manual"`
    Private                 bool    `json:"private"`
    Flagged                 bool    `json:"flagged"`
    Gear_id                 string  `json:"gear_id"`
    From_accepted_tag       bool    `json:"from_accepted_tag"`
    Average_speed           float32 `json:"average_speed"`
    Max_speed               float32 `json:"max_speed"`
    Average_cadence         float32 `json:"average_cadence"`
    Average_temp            float32 `json:"average_temp"`
    Average_watts           float32 `json:"average_watts"`
    Weighted_average_watts  float32 `json:"weighted_average_watts"`
    Kilojoules              float32 `json:"kilojoules"`
    Device_watts            bool    `json:"device_watts"`
    Has_Heartrate           bool    `json:"has_heartrate"`
    Max_watts               float32 `json:"max_watts"`
    Elev_high               float32 `json:"elev_high"`
    Elev_low                float32 `json:"elev_low"`
    Pr_count                int     `json:"pr_count"`
    Total_photo_count       int     `json:"total_photo_count"`
    Has_Kudoed              bool    `json:"has_kudoed"`
    Workout_type            int     `json:"workout_type"`
    Suffer_score            int     `json:"suffer_score"`
    Description             string  `json:"description"`
    Calories                float32 `json:"calories"`
    Segment_efforts         []Segment_Effort `json:"segment_efforts"`
    Splits_metric           []Split `json:"splits_metric"`
    Laps                    []Lap   `json:"laps"`
    Gear                    Gear    `json:"gear"`
    Partner_brand_tag       bool    `json:"partner_brand_tag"` // TODO: qq: not certain
    Photos                  Photos  `json:"photos"`
    Highlighted_kudosers    []Highlighted_Kudoser `json:"highlighted_kudosers"`
    Device_name             string  `json:"device_name"`
    Embed_token             string  `json:"embed_token"`
    Segment_leaderboard_opt_out bool `json:"segment_leaderboard_opt_out"`
    Leaderboard_opt_out     bool    `json:"leaderboard_opt_out"`
}

type Activity_Pair struct {
    Id                      int     `json:"id"`
    Resource_state          int     `json:"resource_state"`
}

type Segment_Effort struct {
    Id                      int     `json:"id"`
    Resource_state          int     `json:"resource_state"`
    Name                    string  `json:"name"`
    Activity                Activity_Pair `json:"activity"`
    Athlete                 Activity_Pair `json:"activity"`
    Elapsed_time            int     `json:"elapsed_time"`
    Moving_time             int     `json:"moving_time"`
    Start_date              string  `json:"start_date"`
    Start_date_local        string  `json:"start_date_local"`
    Distance                float32     `json:"distance"`
    Start_index             int     `json:"start_index"`
    End_index               int     `json:"end_index"`
    Average_cadence         float32     `json:"average_candence"`
    Device_watts            bool    `json:"device_watts"`
    Average_watts           float32     `json:"average_watts"`
    Segment                 Activity_Segment `json:"segment"`
    Kom_rank                int     `json:"kom_rank"` // TODO: qq: these 3 are uncertain
    Pr_rank                 int     `json:"pr_rank"`
    Achievments             []string `json:"achievments"`
    Hidden                  bool    `json:"hidden"`
}

type Activity_Segment struct {
    Id                      int     `json:"id"`
    Resource_state          int     `json:"resource_state"`
    Name                    string  `json:"name"`
    Activity_type           string  `json:"activity_type"`
    Distance                float32 `json:"distance"`
    Average_grade           float32 `json:"average_grade"`
    Maximum_grade           float32 `json:"maximum_grade"`
    Elevation_high          float32 `json:"elevation_high"`
    Elevation_low           float32 `json:"elevation_low"`
    Start_latlng            []float64 `json:"start_latlng"`
    End_latlng              []float64 `json:"end_latlng"`
    Start_latitude          float64 `json:"start_latitude"`
    Start_longitude         float64 `json:"start_longitude"`
    End_latitude            float64 `json:"end_latitude"`
    End_longitude           float64 `json:"end_longitude"`
    Climb_category          int     `json:"climb_category"`
    City                    string  `json:"city"`
    State                   string  `json:"state"`
    Country                 string  `json:"country"`
    Private                 bool    `json:"private"`
    Hazardous               bool    `json:"hazardous"`
    Starred                 bool    `json:"starred"`
}

type Split struct {
    Distance                float32 `json:"distance"`
    Elapsed_time            float32 `json:"elapsed_time"`
    Elevation_difference    float32 `json:"elevation_difference"`
    Moving_time             float32 `json:"moving_time"`
    Split                   int     `json:"split"`
    Average_speed           float32 `json:"average_speed"`
    Pace_zone               int     `json:"pace_zone"`
}

type Lap struct {
    Id                      int     `json:"id"`
    Resource_state          int     `json:"resource_state"`
    Name                    string  `json:"name"`
    Activity                Activity_Pair `json:"activity"`
    Athlete                 Activity_Pair `json:"athlete"`
    Elapsed_time            int     `json:"elapsed_time"`
    Moving_time             int     `json:"moving_time"`
    Start_date              string  `json:"start_date"`
    Start_date_local        string  `json:"start_date_local"`
    Distance                float32 `json:"distance"`
    Start_index             int     `json:"start_index"`
    End_index               int     `json:"end_index"`
    Total_elevation_gain    float32     `json:"total_elevation_gain"`
    Average_speed           float32 `json:"average_speed"`
    Max_speed               float32 `json:"max_speed"`
    Average_cadence         float32 `json:"average_cadence"`
    Device_watts            bool    `json:"device_watts"`
    Average_watts           float32 `json:"average_watts"`
    Lap_index               int     `json:"lap_index"`
    Split                   int     `json:"split"`
}

type Gear struct {
    Id                      string  `json:"id"`
    Primary                 bool    `json:"primary"`
    Name                    string  `json:"name"`
    Resource_state          int     `json:"resource_state"`
    Distance                float32     `json:"distance"`
}

type Highlighted_Kudoser struct {
    Destination_url         string  `json:"destination_url"`
    Display_name            string  `json:"display_name"`
    Avatar_url              string  `json:"avatar_url"`
    Show_name               bool    `json:"show_name"`
}

type Photos struct {
    Primary                 Photo   `json:"primary"`
    Use_primary_photo       bool    `json:"use_primary_photo"`
    Count                   int     `json:"count"`
}

type Photo struct {
    Id                      string  `json:"id"`
    Unique_id               string  `json:"unique_id"`
    Urls                    Urls    `json:"urls"`
    Source                  int     `json:"source"`
}

type Urls struct {
    One                     string  `json:"100"`
    Six                     string  `json:"600"`
}