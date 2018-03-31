package Galio

import (
	"testing"
	"os"
	"github.com/skarajic/galio/regions"
	"fmt"
)

var key = os.Getenv("API_KEY")

/**
Tests the getSummoner method
 */
func TestNew(t *testing.T) {
	galio := New(key, regions.EUW)

	if galio.Region != regions.EUW {
		t.Errorf(fmt.Sprintf("Wrong region received. should be euw1, got %s", galio.Region))
	}
}