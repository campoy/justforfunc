package backoff

import (
	"time"

	"testing"
)

func TestNextBackOffMillis(t *testing.T) {
	new(ZeroBackOff).BackOff()
}

func TestConstantBackOff(t *testing.T) {
	backoff := NewConstant(time.Second)
	if backoff.Interval != time.Second {
		t.Error("invalid interval")
	}
}
