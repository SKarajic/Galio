/*
Galio is a wrapper for Riot API written in Golang with support for the CommunityDragon CDN. It supports many features
such as multi-region rate-limiting, request recursion, support for CommunityDragon and more.

The wrapper has a few subpackages, which are the following:
*/
// - github.com/SKarajic/Galio/entities (this contains all the entities that are generated by the wrapper)
// - github.com/SKarajic/Galio/Regions (contains all the regions which are supported by RiotAPI)
// - github.com/SKarajic/Galio/ratelimiter (contains a RateLimiter used by Galio)
/*
There are a few more subpackages, however, these are meant for internal use. If you have any use for the subpackages
on their own, feel free to use them!

The source code can be found on https://github.com/SKarajic/Galio. Feel free to create a pull request and help out!

Galio isn’t endorsed by Riot Games and doesn’t reflect the views or opinions of Riot Games or anyone officially
involved in producing or managing League of Legends. League of Legends and Riot Games are trademarks or registered
trademarks of Riot Games, Inc. League of Legends © Riot Games, Inc.
*/
package Galio

import (
	"github.com/skarajic/galio/entities"
	"github.com/skarajic/galio/handlers"
	"github.com/skarajic/galio/regions"
)

type Galio entities.Galio
type SummonerInput entities.SummonerInput

// Creates a new Galio wrapper entity, handling all the requests to the API
//
// parameters:
// - apiKey (string) : your API key
// - region (Region) : the region of choice (found in the github.com/SKarajic/Galio/regions package)
//
// returns: (Galio) the Galio wrapper object
func New(ApiKey string, Region regions.Region) Galio {
	return Galio{ApiKey: ApiKey, Region: Region}
}

// Returns a new Summoner entity, containing all data of a summoner from a specific region
//
// parameters:
// - input (SummonerInput) : a SummonerInput entity containing either a summoner ID, summoner name or account ID
//
// returns: (Summoner) returns a new Summoner Object

func (g *Galio) GetSummoner(input SummonerInput) entities.Summoner {
	return handlers.GetSummoner(entities.Galio(*g), entities.SummonerInput(input))
}

// Returns a new MatchList entity, containing the summary of matches of a specific summoner
//
// parameters:
// - input (SummonerInput) : a SummonerInput entity containing either a summoner ID, summoner name or account ID
// - recent (boolean)      : a boolean to decide if the last 20 matches should be returned or all matches
//
// returns: (MatchList) returns a new MatchList Object filled with MatchSummary objects
func (g *Galio) GetMatchList(input SummonerInput, recent bool) entities.MatchList {
	return handlers.GetMatchList(entities.Galio(*g), entities.SummonerInput(input), recent)
}

// Returns a new MatchList entity, containing the data of a specific matches
//
// parameters:
// - matchId (uint64) : the ID of a match which you can obtain through the matchlist
//
// returns: (MatchList) returns a new MatchList Object filled with MatchSummary objects
func (g *Galio) GetMatch(matchId uint64) entities.Match {
	return handlers.GetMatch(entities.Galio(*g), matchId)
}
