package queues

import (
	"fmt"
	"testing"
)

func TestQueues(t *testing.T) {
	queue := Solo
	queue = Flex

	if queue != Flex {
		fmt.Print("Wrong queue\n")
	}
}
