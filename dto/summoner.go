package dto

// SummonerDTO SummonerDTO
type SummonerDTO struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Level        uint   `json:"summonerLevel"`
	IconID       uint   `json:"profileIconId"`
	AccountID    uint64 `json:"accountId"`
	RevisionDate uint64 `json:"revisionDate"`
}
