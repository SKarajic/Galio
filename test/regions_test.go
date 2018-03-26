package test

import (
	"testing"
	"../regions"
	"fmt"
)

func TestRegions(t *testing.T) {
	reg := regions.EUNE
	reg = regions.EUW
	reg = regions.LAN
	reg = regions.LAS
	reg = regions.OCE
	reg = regions.TR
	reg = regions.NA
	reg = regions.BR
	reg = regions.JP
	reg = regions.KR
	reg = regions.RU
	fmt.Sprint(reg)
}