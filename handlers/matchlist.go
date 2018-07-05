package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/skarajic/galio/dto"
)

// GetMatchList returns a new MatchListDTO entity, containing the summary of matches of a specific summoner
//
// parameters:
// - wrapper (Galio)       : the Wrapper
// - input (SummonerInput) : a SummonerInput entity containing either a summoner ID, summoner name or account ID
// - recent (boolean)      : a boolean to decide if the last 20 matches should be returned or all matches
//
// returns: (MatchListDTO) returns a new MatchListDTO Object filled with MatchReferenceDTO objects
func GetMatchList(wrapper dto.Galio, input dto.SummonerInput, recent bool) dto.MatchListDTO {
	var data string
	var accountID uint64
	var matchlist dto.MatchListDTO
	matchlistEndpoint := "match/v3/matchlists/by-account/"

	if input.AccountID != 0 {
		accountID = input.AccountID
	} else {
		accountID = GetSummoner(wrapper, input).AccountID
	}

	matchlistEndpoint += fmt.Sprint(accountID)

	if recent {
		matchlistEndpoint += "/recent"
	}

	data = wrapper.GetRiotData(matchlistEndpoint)
	json.Unmarshal([]byte(data), &matchlist)
	return matchlist
}
