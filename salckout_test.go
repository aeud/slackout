package slackout

import (
	"fmt"
	"testing"
)

func TestW(t *testing.T) {
	fmt.Fprint(W, "test")
}
