package entities

import (
	"testing"
	"os"
	"github.com/skarajic/galio/regions"
	"fmt"
)

var key = os.Getenv("API_KEY")

func TestWrapper(t *testing.T) {
	galio := Galio{key, regions.EUW}

	if galio.Region != regions.EUW {
		t.Errorf(fmt.Sprintf("Wrong region received. should be euw1, got %s", galio.Region))
	}
}