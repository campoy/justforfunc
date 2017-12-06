package backoff

import (
	"testing"
	"time"
)

func TestLinear(t *testing.T) {
	bmult := NewLinear(time.Minute, 10*time.Minute, 0, 4)

	bmult.increment()

	if bmult.currentInterval != time.Minute*4 {
		t.Errorf("increment did not work got %s expected %s.", bmult.currentInterval, time.Minute*2)
	}

	bincr := NewLinear(time.Minute, 10*time.Minute, time.Minute, 1)

	bincr.increment()

	if bincr.currentInterval != time.Minute*2 {
		t.Errorf("increment did not work got %s expected %s.", bincr.currentInterval, time.Minute*2)
	}

	bmultincr := NewLinear(time.Minute, 10*time.Minute, time.Minute, 2)
	bmultincr.increment()
	if bmultincr.currentInterval != time.Minute*3 {
		t.Errorf("increment did not work got %s expected %s.", bmultincr.currentInterval, time.Minute*3)
	}

	bmultincr.Reset()

	if bmultincr.currentInterval != time.Minute {
		t.Errorf("reset did not work got %s expected %s.", bmultincr.currentInterval, time.Minute)
	}
}
