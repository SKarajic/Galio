package galio

import (
	"fmt"
	"./entities"
	"./regions"
	"./urls"
	"./util"
)

type Galio struct {
	ApiKey string
	Region regions.Region
}

func New(ApiKey string, Region regions.Region) Galio {
	return Galio{ApiKey, Region}
}


func (g *Galio) getData(endpoint string) string {
	url := fmt.Sprintf(urls.API, g.Region, endpoint)

	m := make(map[string] string)
	m["X-Riot-Token"] = g.ApiKey

	return util.GetJSON(url, m)
}

func (g *Galio) GetChampions() []entities.Champion {
	data := g.getData("platform/v3/champions")
	fmt.Println(data)
	return nil
}
