package queues

import (
	"fmt"
	"testing"
)

func TestQueues(t *testing.T) {
	reg := Solo
	reg = Flex
	fmt.Sprint(reg)
}
