package df

import (
	"testing"
	"time"
)

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero_delay", 0 * time.Second},
		{"quarter-of-a-second_delay", 250 * time.Millisecond},
		{"half-second_delay", 500 * time.Millisecond},
	}

	for _, e := range theTests {
		orderOfLeaving = []string{}

		eatTime = 0 * e.delay
		thinkTime = 0 * e.delay
		sleepTime = 0 * e.delay

		dine()
		if len(orderOfLeaving) != len(philosophers) {
			t.Errorf("%s: incorrect length of slice. Expected %d but got %d!\n", e.name, len(philosophers), len(orderOfLeaving))
		}
	}
}
