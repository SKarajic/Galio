package handlers

import (
	"fmt"
	"github.com/skarajic/galio/entities"
	"encoding/json"
)

func GetSummoner(wrapper entities.Galio, input entities.SummonerInput) entities.Summoner {
	var data string
	var summoner entities.Summoner
	endpoint := "summoner/v3/summoners/"

	switch {
	case input.AccountId != 0:
		data = wrapper.GetData(endpoint + "by-account/" + fmt.Sprint(input.AccountId))
	case input.SummonerName != "":
		data = wrapper.GetData(endpoint + "by-name/" + input.SummonerName)
	case input.SummonerId != 0:
		data = wrapper.GetData(endpoint + fmt.Sprint(input.SummonerId))
	}

	json.Unmarshal([]byte(data), &summoner)
	return summoner
}