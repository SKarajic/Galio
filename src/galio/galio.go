/*
The Wrapper Package
 */
package galio

import (
	"./handlers"
	"./entities"
	"./regions"
)

type Galio entities.Galio
type SummonerInput entities.SummonerInput

/*
Generates a new Galio Wrapper Object
 */
func New(ApiKey string, Region regions.Region) Galio {
	return Galio{ApiKey: ApiKey, Region: Region}
}

func (g *Galio) GetSummoner(input SummonerInput) entities.Summoner {
	return handlers.GetSummoner(entities.Galio(*g), entities.SummonerInput(input))
}
