package Galio

import (
	"github.com/skarajic/galio/regions"
	"github.com/skarajic/galio/entities"
	"testing"
	"sync"
	"os"
)

var key = os.Getenv("API_KEY")
var galio = New(key, regions.EUW)
var wg = sync.WaitGroup{}

/**
Tests the getSummoner method
 */
func TestGalio_GetSummoner(t *testing.T) {
	var sum1, sum2, sum3 entities.Summoner

	si1 := SummonerInput{ SummonerName: "IAmTheWhite" }
	si2	:= SummonerInput{ SummonerId: 69658457 }
	si3 := SummonerInput{ AccountId: 219406964 }

	wg.Add(3)
	go func() {
		defer wg.Done()
		sum1 = galio.GetSummoner(si1)
	}()
	go func() {
		defer wg.Done()
		sum2 = galio.GetSummoner(si2)

	}()
	go func() {
		defer wg.Done()
		sum3 = galio.GetSummoner(si3)
	}()
	wg.Wait()

	if sum1.Id != 69658457 || sum2.Id != 69658457 || sum3.Id != 69658457 {
		t.Errorf("Wrong summonerID received")
	}
}

/**
Test the getMatchList method
 */
func TestGalio_GetMatchList(t *testing.T) {
	var ml1, ml2, ml3, ml4 entities.MatchList

	si1 := SummonerInput{ SummonerName: "IAmTheWhite" }
	si2	:= SummonerInput{ SummonerId: 69658457 }
	si3 := SummonerInput{ AccountId: 219406964 }

	wg.Add(4)
	go func() {
		defer wg.Done()
		ml1 = galio.GetMatchList(si1, false)
	}()
	go func() {
		defer wg.Done()
		ml2 = galio.GetMatchList(si2, false)
	}()
	go func() {
		defer wg.Done()
		ml3 = galio.GetMatchList(si3, false)
	}()
	go func() {
		defer wg.Done()
		ml4 = galio.GetMatchList(si3, true)
	}()
	wg.Wait()

	if len(ml1.Matches) < 0 || len(ml2.Matches) < 0 || len(ml3.Matches) < 0 || len(ml4.Matches) < 0 {
		t.Errorf("Not the correct matchlist data received")
	}
}