package handlers

import (
	"fmt"
	"github.com/skarajic/galio/entities"
	"encoding/json"
)

func GetMatchList(wrapper entities.Galio, input entities.SummonerInput, recent bool) entities.MatchList {
	var data string
	var accountId uint64
	var matchlist entities.MatchList
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