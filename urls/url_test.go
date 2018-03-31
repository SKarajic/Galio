package urls

import (
	"testing"
	"fmt"
)

func TestUrls(t *testing.T) {
	url := API
	fmt.Sprintf(url, "euw1", "test")
}