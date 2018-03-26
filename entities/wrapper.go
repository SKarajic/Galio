package entities

import (
	"github.com/skarajic/galio/regions"
	"github.com/skarajic/galio/urls"
	"github.com/skarajic/galio/util"
	"fmt"
)

type SummonerInput struct {
	SummonerName string
	SummonerId uint64
	AccountId uint64
}

type Galio struct {
	ApiKey string
	Region regions.Region
}


/*
Handles requests to Riot API
 */
func (g *Galio) GetData(endpoint string) string {
	url := fmt.Sprintf(urls.API, g.Region, endpoint)

	m := make(map[string] string)
	m["X-Riot-Token"] = g.ApiKey

	return util.GetJSON(url, m)
}
