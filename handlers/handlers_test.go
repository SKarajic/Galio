package handlers

import (
	"fmt"
	"github.com/skarajic/galio/dto"
	"github.com/skarajic/galio/regions"
	"os"
	"sync"
	"testing"
)

var key = os.Getenv("API_KEY")
var wrapper = dto.Galio{key, regions.EUW}
var wg = sync.WaitGroup{}

/**
Tests the getSummoner method
*/
func TestGetSummoner(t *testing.T) {
	var sums [3]dto.SummonerDTO
	var sis [3]dto.SummonerInput

	sis[0] = dto.SummonerInput{SummonerName: "IAmTheWhite"}
	sis[1] = dto.SummonerInput{SummonerId: 69658457}
	sis[2] = dto.SummonerInput{AccountId: 219406964}

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
		if sum.Id != 69658457 {
			t.Errorf(fmt.Sprintf("Wrong summonerID received. should be 69658457, got %d", sum.Id))
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
	sis[1] = dto.SummonerInput{SummonerId: 69658457}
	sis[2] = dto.SummonerInput{AccountId: 219406964}

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
	si := dto.SummonerInput{AccountId: 219406964}
	ml := GetMatchList(wrapper, si, false)

	mId := ml.Matches[0].MatchId
	match := GetMatch(wrapper, mId)

	if match.MatchId != mId {
		t.Errorf("Wrong matchId received. should be %d, got %d", mId, match.MatchId)
	}
}
