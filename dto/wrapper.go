package dto

import (
	"fmt"
	"github.com/skarajic/galio/maps"
	"github.com/skarajic/galio/queues"
	"github.com/skarajic/galio/regions"
	"github.com/skarajic/galio/urls"
	"github.com/skarajic/galio/util"
)

type QueueInput struct {
	Map   maps.Map
	Queue queues.QueueType
}

type SummonerInput struct {
	SummonerName string
	SummonerId   uint64
	AccountId    uint64
}

type Galio struct {
	ApiKey string
	Region regions.Region
}

/*
Handles requests to Riot API
*/
func (g *Galio) GetRiotData(endpoint string) string {
	url := fmt.Sprintf(urls.API, g.Region, endpoint)

	m := make(map[string]string)
	m["X-Riot-Token"] = g.ApiKey

	return util.GetJSON(url, m)
}
