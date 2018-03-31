package handlers

import (
	"github.com/skarajic/galio/regions"
	"github.com/skarajic/galio/entities"
	"testing"
	"sync"
	"os"
	"fmt"
)

var key = os.Getenv("API_KEY")
var wrapper = entities.Galio{key, regions.EUW}
var wg = sync.WaitGroup{}

/**
Tests the getSummoner method
 */
func TestGetSummoner(t *testing.T) {
	var sums [3]entities.Summoner

	si1 := entities.SummonerInput{ SummonerName: "IAmTheWhite" }
	si2	:= entities.SummonerInput{ SummonerId: 69658457 }
	si3 := entities.SummonerInput{ AccountId: 219406964 }

	wg.Add(3)
	go func() {
		defer wg.Done()
		sums[0] = GetSummoner(wrapper, si1)
	}()
	go func() {
		defer wg.Done()
		sums[1] = GetSummoner(wrapper, si2)

	}()
	go func() {
		defer wg.Done()
		sums[2] = GetSummoner(wrapper, si3)
	}()
	wg.Wait()

	for i := 0; i < 3; i++ {
		if sums[i].Id != 69658457 {
			t.Errorf(fmt.Sprintf("Wrong summonerID received. should be 69658457, got %d", sums[i].Id))
		}
	}
}

/**
Test the getMatchList method
 */
func TestGetMatchList(t *testing.T) {
	var ml1, ml2, ml3, ml4 entities.MatchList

	si1 := entities.SummonerInput{ SummonerName: "IAmTheWhite" }
	si2	:= entities.SummonerInput{ SummonerId: 69658457 }
	si3 := entities.SummonerInput{ AccountId: 219406964 }

	wg.Add(4)
	go func() {
		defer wg.Done()
		ml1 = GetMatchList(wrapper, si1, false)
	}()
	go func() {
		defer wg.Done()
		ml2 = GetMatchList(wrapper, si2, false)
	}()
	go func() {
		defer wg.Done()
		ml3 = GetMatchList(wrapper, si3, false)
	}()
	go func() {
		defer wg.Done()
		ml4 = GetMatchList(wrapper, si3, true)

	}()
	wg.Wait()

	if len(ml1.Matches) < 0 || len(ml2.Matches) < 0 || len(ml3.Matches) < 0 || len(ml4.Matches) < 0 {
		t.Errorf("Not the correct matchlist data received")
	}
}