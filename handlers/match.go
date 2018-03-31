package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/skarajic/galio/entities"
)

func GetMatch(wrapper entities.Galio, matchId uint64) entities.Match {
	var data string
	var match entities.Match
	matchlistEndpoint := "match/v3/matches/" + fmt.Sprint(matchId)

	data = wrapper.GetRiotData(matchlistEndpoint)
	json.Unmarshal([]byte(data), &match)
	return match
}
