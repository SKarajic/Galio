package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/skarajic/galio/dto"
)

// GetSummoner returns a new SummonerDTO entity, containing all data of a summoner from a specific region
//
// parameters:
// - wrapper (Galio)       : the Wrapper
// - input (SummonerInput) : a SummonerInput entity containing either a summoner ID, summoner name or account ID
//
// returns: (SummonerDTO) returns a new SummonerDTO Object
func GetSummoner(wrapper dto.Galio, input dto.SummonerInput) dto.SummonerDTO {
	var data string
	var summoner dto.SummonerDTO
	endpoint := "summoner/v3/summoners/"

	switch {
	case input.AccountID != 0:
		data = wrapper.GetRiotData(endpoint + "by-account/" + fmt.Sprint(input.AccountID))
	case input.SummonerName != "":
		data = wrapper.GetRiotData(endpoint + "by-name/" + input.SummonerName)
	case input.SummonerID != 0:
		data = wrapper.GetRiotData(endpoint + fmt.Sprint(input.SummonerID))
	}

	json.Unmarshal([]byte(data), &summoner)
	return summoner
}
