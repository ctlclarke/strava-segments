package response_bodies

type Club struct {	
	Id 						int `json:"id"`
	Resource_state			int `json:"resource_state"`
	Name					string `json:"name"`
	Profile_medium			string `json:"profile_medium"`
	Profile 				string `json:"profile"`
	Cover_photo 			string `json:"cover_photo"`
	Cover_photo_small		string `json:"cover_photo_small"`
	Sport_type				string `json:"sport_type"`
	City 					string `json:"city"`
	State 					string `json:"state"`
	Country 				string `json:"country"`
	Private					bool `json:"private"`
	Nember_count 			int `json:"member_count"`
	Featured 				bool `json:"featured"`
	Verified 				bool `json:"verified"`
	Url 					string `json:"url"`
	Membership 				string `json:"membership"`
	Admin					bool `json:"admin"`
	Owner 					bool `json:"owner"`
	Description 			string `json:"description"`
	Club_type 				string `json:"club_type"`
	Following_count 		int `json:"following_count"`
}