package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/skarajic/galio/dto"
)

func GetMatchList(wrapper dto.Galio, input dto.SummonerInput, recent bool) dto.MatchListDTO {
	var data string
	var accountId uint64
	var matchlist dto.MatchListDTO
	matchlistEndpoint := "match/v3/matchlists/by-account/"

	if input.AccountId != 0 {
		accountId = input.AccountId
	} else {
		accountId = GetSummoner(wrapper, input).AccountId
	}

	matchlistEndpoint += fmt.Sprint(accountId)

	if recent {
		matchlistEndpoint += "/recent"
	}

	data = wrapper.GetRiotData(matchlistEndpoint)
	json.Unmarshal([]byte(data), &matchlist)
	return matchlist
}
