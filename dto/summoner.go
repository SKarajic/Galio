package dto

type SummonerDTO struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Level        uint   `json:"summonerLevel"`
	IconId       uint   `json:"profileIconId"`
	AccountId    uint64 `json:"accountId"`
	RevisionDate uint64 `json:"revisionDate"`
}
