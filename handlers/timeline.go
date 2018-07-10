package handlers

import (
	"fmt"

	"github.com/skarajic/galio/dto"
)

// GetTimeLine returns a new TimeLineDTO entity, containing the timeline of a specific match.
// Handles request:  /lol/match/v3/timelines/by-match/{matchId}
//
// parameters:
// - wrapper (Galio)       : the Wrapper
// - matchID (uint64) : the ID of a match which you can obtain through the match
//
// returns: (TimelineDTO) returns a new MatchDTO Object
func GetTimeLine(wrapper dto.Galio, matchID uint64) dto.TimeLineDTO {
	var data string
	timeLineEndpoint := "match/v3/timelines/by-match/" + fmt.Sprint(matchID)

	data = wrapper.GetRiotData(timeLineEndpoint)
	return dto.NewTimeLine(data)
}
