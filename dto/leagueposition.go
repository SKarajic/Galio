package dto

// LeaguePositionDTO LeaguePositionDTO
type LeaguePositionDTO struct {
	QueueType  string        `json:"queueType"`
	Tier       string        `json:"tier"`
	Rank       string        `json:"rank"`
	LP         int           `json:"leaguePoints"`
	PlayerID   string        `json:"playerOrTeamId"`
	PlayerName string        `json:"playerOrTeamName"`
	LeagueID   string        `json:"leagueId"`
	LeagueName string        `json:"leagueName"`
	Wins       uint          `json:"wins"`
	Losses     uint          `json:"losses"`
	MiniSeries MiniSeriesDTO `json:"miniSeries"`
	HotStreak  bool          `json:"hotStreak"`
	Veteran    bool          `json:"veteran"`
	Inactive   bool          `json:"inactive"`
	FreshBlood bool          `json:"freshBlood"`
}

// MiniSeriesDTO MiniSeriesDTO
type MiniSeriesDTO struct {
	Wins     uint8  `json:"wins"`
	Losses   uint8  `json:"losses"`
	Target   uint8  `json:"target"`
	Progress string `json:"progress"`
}
