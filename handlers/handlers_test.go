package handlers

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/skarajic/galio/dto"
	"github.com/skarajic/galio/regions"
)

var key = os.Getenv("API_KEY")
var wrapper = dto.Galio{
	APIKey: key,
	Region: regions.EUW,
}
var wg = sync.WaitGroup{}

/**
Tests the getSummoner method
*/
func TestGetSummoner(t *testing.T) {
	var sums [3]dto.SummonerDTO
	var sis [3]dto.SummonerInput

	sis[0] = dto.SummonerInput{SummonerName: "IAmTheWhite"}
	sis[1] = dto.SummonerInput{SummonerID: 69658457}
	sis[2] = dto.SummonerInput{AccountID: 219406964}

	wg.Add(3)
	for i := range sis {
		go func(i int) {
			defer wg.Done()
			sums[i] = GetSummoner(wrapper, sis[i])
		}(i)
	}
	wg.Wait()

	fmt.Println(sums[0])

	for _, sum := range sums {
		if sum.ID != 69658457 {
			t.Errorf(fmt.Sprintf("Wrong summonerID received. should be 69658457, got %d", sum.ID))
		}
	}
}

/**
Test the getMatchList method
*/
func TestGetMatchList(t *testing.T) {
	var mls [4]dto.MatchListDTO
	var sis [3]dto.SummonerInput

	sis[0] = dto.SummonerInput{SummonerName: "IAmTheWhite"}
	sis[1] = dto.SummonerInput{SummonerID: 69658457}
	sis[2] = dto.SummonerInput{AccountID: 219406964}

	wg.Add(4)
	for i := range sis {
		go func(i int) {
			defer wg.Done()
			mls[i] = GetMatchList(wrapper, sis[i], false)
		}(i)
	}
	go func() {
		defer wg.Done()
		mls[3] = GetMatchList(wrapper, sis[2], true)
	}()
	wg.Wait()

	for i, ml := range mls {
		if len(ml.Matches) < 0 {
			t.Errorf("Not the correct matchlist data received from matchlist %d", i)
		}
	}
}

/**
Test the getMatch method
*/
func TestGetMatch(t *testing.T) {
	si := dto.SummonerInput{AccountID: 219406964}
	ml := GetMatchList(wrapper, si, false)

	mID := ml.Matches[0].MatchID
	match := GetMatch(wrapper, mID)

	if match.MatchID != mID {
		t.Errorf("Wrong matchId received. should be %d, got %d", mID, match.MatchID)
	}
}
