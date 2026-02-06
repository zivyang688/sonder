package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ziv/sonder/tools/ter"
	test_tools "github.com/ziv/sonder/tools/test-tools"
)

func TestTimeTracker(t *testing.T) {
	test_tools.ElaspedTimeTrack(func() {
		time.Sleep(1 * time.Second)
	})
}

func TestTer(t *testing.T) {
	rst := ter.Ter(1 > 3, "1>3", "1<3")
	fmt.Printf("rst: %v\n", rst)
}
