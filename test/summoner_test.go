package test

import (
	"github.com/skarajic/galio/regions"
	"github.com/skarajic/galio"
	"os"
	"testing"
)



func Test(t *testing.T) {
	key := os.Getenv("API_KEY")
	Galio := galio.New(key, regions.EUW)

	si1 := galio.SummonerInput{ SummonerName: "IAmTheWhite"}
	si2	:= galio.SummonerInput{ SummonerId: 69658457 }
	si3 := galio.SummonerInput{ AccountId: 219406964 }

	sum1 := Galio.GetSummoner(si1)
	sum2 := Galio.GetSummoner(si2)
	sum3 := Galio.GetSummoner(si3)

	if sum1.Id != 69658457 || sum2.Id != 69658457 || sum3.Id != 69658457 {
		t.Errorf("Wrong summonerID received")
	}
}

