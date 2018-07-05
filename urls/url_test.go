package urls

import (
	"fmt"
	"testing"
)

func TestUrls(t *testing.T) {
	url := fmt.Sprintf(API, "euw1", "test")

	if url != "https://euw1.api.riotgames.com/lol/test" {
		fmt.Printf("Incorrect URL, got returned: %s\n", url)
	}
}
