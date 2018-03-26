/*
The Wrapper Package
 */
package Galio

import (
	"github.com/skarajic/galio/handlers"
	"github.com/skarajic/galio/entities"
	"github.com/skarajic/galio/regions"
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
