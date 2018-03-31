/*
The Wrapper Package
*/
package Galio

import (
	"github.com/skarajic/galio/entities"
	"github.com/skarajic/galio/handlers"
	"github.com/skarajic/galio/regions"
)

type Galio entities.Galio
type SummonerInput entities.SummonerInput

/**
Generates a new Galio Wrapper Object
*/
func New(ApiKey string, Region regions.Region) Galio {
	return Galio{ApiKey: ApiKey, Region: Region}
}

/**
Gets a summoner
*/
func (g *Galio) GetSummoner(input SummonerInput) entities.Summoner {
	return handlers.GetSummoner(entities.Galio(*g), entities.SummonerInput(input))
}

/**
Gets a matchlist
*/
func (g *Galio) GetMatchList(input SummonerInput, recent bool) entities.MatchList {
	return handlers.GetMatchList(entities.Galio(*g), entities.SummonerInput(input), recent)
}

/**
Gets a match
*/
func (g *Galio) GetMatch(matchId uint64) entities.Match {
	return handlers.GetMatch(entities.Galio(*g), matchId)
}
