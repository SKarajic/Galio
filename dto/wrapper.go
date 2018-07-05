package dto

import (
	"fmt"

	"github.com/skarajic/galio/maps"
	"github.com/skarajic/galio/queues"
	"github.com/skarajic/galio/regions"
	"github.com/skarajic/galio/urls"
	"github.com/skarajic/galio/util"
)

// QueueInput is a Queue Input Entity
type QueueInput struct {
	Map   maps.Map
	Queue queues.QueueType
}

// SummonerInput is a Summoner Input Entity
type SummonerInput struct {
	SummonerName string
	SummonerID   uint64
	AccountID    uint64
}

// Galio is a wrapper for RiotAPI
type Galio struct {
	APIKey string
	Region regions.Region
}

// GetRiotData is a request handler for RiotAPI
//
// parameters:
// - endpoint (string) : the Endpoint from RiotAPI
//
// returns: (string) returns a JSON from a RiotAPI endpoint
func (g *Galio) GetRiotData(endpoint string) string {
	url := fmt.Sprintf(urls.API, g.Region, endpoint)

	m := make(map[string]string)
	m["X-Riot-Token"] = g.APIKey

	return util.GetJSON(url, m)
}
