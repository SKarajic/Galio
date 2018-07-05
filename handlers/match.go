package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/skarajic/galio/dto"
)

// GetMatch returns a new MatchListDTO entity, containing the data of a specific matches
//
// parameters:
// - wrapper (Galio)       : the Wrapper
// - matchID (uint64) : the ID of a match which you can obtain through the matchlist
//
// returns: (MatchListDTO) returns a new MatchListDTO Object filled with MatchReferenceDTO objects
func GetMatch(wrapper dto.Galio, matchID uint64) dto.MatchDTO {
	var data string
	var match dto.MatchDTO
	matchlistEndpoint := "match/v3/matches/" + fmt.Sprint(matchID)

	data = wrapper.GetRiotData(matchlistEndpoint)
	json.Unmarshal([]byte(data), &match)
	return match
}
