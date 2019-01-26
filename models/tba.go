package models

type Match []struct {
	ActualTime int64 `json:"actual_time"`
	Alliances  struct {
		Blue Alliance `json:"blue"`
		Red  Alliance `json:"red"`
	} `json:"alliances"`
	CompLevel       string `json:"comp_level"`
	EventKey        string `json:"event_key"`
	Key             string `json:"key"`
	MatchNumber     int64  `json:"match_number"`
	PredictedTime   int64  `json:"predicted_time"`
	SetNumber       int64  `json:"set_number"`
	Time            int64  `json:"time"`
	WinningAlliance string `json:"winning_alliance"`
}

type Alliance struct {
	DQTeamKeys        []string `json:"dq_team_keys"`
	Score             int64    `json:"score"`
	SurrogateTeamKeys []string `json:"surrogate_team_keys"`
	TeamKeys          []string `json:"team_keys"`
}

type Error struct {
	Errors []map[string]string `json:"Errors"`
}
