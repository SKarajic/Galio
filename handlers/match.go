package handlers

import (
	"fmt"

	"github.com/skarajic/galio/dto"
)

// GetMatch returns a new MatchListDTO entity, containing the data of a specific matches.
// Handles request:  /lol/match/v3/matches/{matchId}
//
// parameters:
// - wrapper (Galio)       : the Wrapper
// - matchID (uint64) : the ID of a match which you can obtain through the match
//
// returns: (MatchDTO) returns a new MatchDTO Object
func GetMatch(wrapper dto.Galio, matchID uint64) dto.MatchDTO {
	var data string
	matchlistEndpoint := "match/v3/matches/" + fmt.Sprint(matchID)

	data = wrapper.GetRiotData(matchlistEndpoint)
	return dto.NewMatch(data)
}
