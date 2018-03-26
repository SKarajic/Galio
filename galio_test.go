package Galio

import (
	"./regions"
	"os"
	"testing"
)

func Test(t *testing.T) {
	key := os.Getenv("API_KEY")
	galio := New(key, regions.EUW)

	si1 := SummonerInput{ SummonerName: "IAmTheWhite" }
	si2	:= SummonerInput{ SummonerId: 69658457 }
	si3 := SummonerInput{ AccountId: 219406964 }

	sum1 := galio.GetSummoner(si1)
	sum2 := galio.GetSummoner(si2)
	sum3 := galio.GetSummoner(si3)

	if sum1.Id != 69658457 || sum2.Id != 69658457 || sum3.Id != 69658457 {
		t.Errorf("Wrong summonerID received")
	}
}

