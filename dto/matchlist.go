package dto

type MatchListDTO struct {
	Matches    []MatchReferenceDTO `json:"matches"`
	EndIndex   uint64              `json:"endIndex"`
	StartIndex uint64              `json:"startIndex"`
	TotalGames uint64              `json:"totalGames"`
}

type MatchReferenceDTO struct {
	Lane       string `json:"lane"`
	MatchId    uint64 `json:"gameId"`
	Champion   uint64 `json:"champion"`
	PlatformID string `json:"platformId"`
	Timestamp  uint64 `json:"timestamp"`
	Queue      uint64 `json:"queue"`
	Role       string `json:"role"`
	Season     uint64 `json:"season"`
}
