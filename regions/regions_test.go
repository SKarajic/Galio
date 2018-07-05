package regions

import (
	"fmt"
	"testing"
)

func TestRegions(t *testing.T) {
	reg := EUNE
	reg = EUW
	reg = LAN
	reg = LAS
	reg = OCE
	reg = TR
	reg = NA
	reg = BR
	reg = JP
	reg = KR
	reg = RU

	if reg != RU {
		fmt.Print("Wrong region\n")
	}
}
