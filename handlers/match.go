package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/skarajic/galio/dto"
)

func GetMatch(wrapper dto.Galio, matchId uint64) dto.MatchDTO {
	var data string
	var match dto.MatchDTO
	matchlistEndpoint := "match/v3/matches/" + fmt.Sprint(matchId)

	data = wrapper.GetRiotData(matchlistEndpoint)
	json.Unmarshal([]byte(data), &match)
	return match
}
