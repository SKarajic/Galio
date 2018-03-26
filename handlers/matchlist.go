package handlers

import (
	"fmt"
	"github.com/skarajic/galio/entities"
	"encoding/json"
)

func GetMatchList(wrapper entities.Galio, input entities.SummonerInput, recent bool) entities.MatchList {
	var summonerData string
	var matchlistData string

	var accountId uint64
	var summoner entities.Summoner
	var matchlist entities.MatchList
	summonerEndpoint := "summoner/v3/summoners/"
	matchlistEndpoint := "match/v3/matchlists/by-account/"

	switch {
	case input.AccountId != 0:
		accountId = input.AccountId

	case input.SummonerName != "":
		summonerData = wrapper.GetRiotData(summonerEndpoint + "by-name/" + input.SummonerName)
		json.Unmarshal([]byte(summonerData), &summoner)
		accountId = summoner.AccountId

	case input.SummonerId != 0:
		summonerData = wrapper.GetRiotData(summonerEndpoint + fmt.Sprint(input.SummonerId))
		json.Unmarshal([]byte(summonerData), &summoner)
		accountId = summoner.AccountId
	}

	matchlistEndpoint += fmt.Sprint(accountId)

	if recent {
		matchlistEndpoint += "/recent"
	}

	matchlistData = wrapper.GetRiotData(matchlistEndpoint)
	json.Unmarshal([]byte(matchlistData), &matchlist)
	return matchlist
}