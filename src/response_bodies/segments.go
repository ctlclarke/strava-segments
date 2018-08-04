package response_bodies

type Segment struct {	
	Id 						int 	`json:"id"`
	Resource_state			int 	`json:"resource_state"`
	Name					string 	`json:"name"`
	Activity_type			string 	`json:"activity_type"`
	Distance 				float32 `json:"distance"`
	Average_grade 			float32 `json:"average_grade"`
	Maximum_grade			float32 `json:"maximum_grade"`
	Elevation_high			float32 `json:"elevation_high"`
	Elevation_low 			float32 `json:"elevation_low"`
	Start_latlng 			[]float64	`json:"start_latlng"`
	Start_latitude 			float64 `json:"start_latitude"`
	Start_longitude			float64 `json:"start_longitude"`
	End_latitude 			float64 `json:"end_latitude"`
	End_longitude 			float64 `json:"end_longitude"`
	Climb_category 			int 	`json:"climb_category"`
	City 					string 	`json:"city"`
	State 					string 	`json:"state"`
	Country					string 	`json:"country"`
	Private 				bool 	`json:"private"`
	Hazardous 				bool 	`json:"hazardous"`
	Starred 				bool 	`json:"starred"`
	Created_at 				string 	`json:"created_at"`
	Updated_at				string 	`json:"updated_at"`
	Total_elevation_gain	float32 `json:"total_elevation_gain"`
	Map						Map 	`json:"map"`
	Effort_count			int 	`json:"effort_count"`
	Athlete_count			int 	`json:"athlete_count"`
	Star_count				int 	`json:"star_count"`
	Athlete_segment_stats	SegStat	`json:"athlete_segment_stats"`
}

type Map struct {
	Id 						string	`json:"id"`
	Polyline				string 	`json:"polyline"`
	Resource_state 			int 	`json:"resource_state"`
	Summary_polyline		string	`json:"summary_polyline"`
}

type SegStat struct {
	Pr_elapsed_time			int		`json:"pr_elapsed_time"`
	Pr_date 				string	`json:"pr_date"`
	Effort_count 			int 	`json:"effort_count"`
}