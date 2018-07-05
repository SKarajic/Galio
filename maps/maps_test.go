package maps

import (
	"fmt"
	"testing"
)

func TestMaps(t *testing.T) {
	playableMap := SummonersRift
	playableMap = TwistedTreeline

	if playableMap != TwistedTreeline {
		fmt.Print("Wrong map\n")
	}
}
