package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/skarajic/galio/dto"
)

func GetSummoner(wrapper dto.Galio, input dto.SummonerInput) dto.SummonerDTO {
	var data string
	var summoner dto.SummonerDTO
	endpoint := "summoner/v3/summoners/"

	switch {
	case input.AccountId != 0:
		data = wrapper.GetRiotData(endpoint + "by-account/" + fmt.Sprint(input.AccountId))
	case input.SummonerName != "":
		data = wrapper.GetRiotData(endpoint + "by-name/" + input.SummonerName)
	case input.SummonerId != 0:
		data = wrapper.GetRiotData(endpoint + fmt.Sprint(input.SummonerId))
	}

	json.Unmarshal([]byte(data), &summoner)
	return summoner
}
