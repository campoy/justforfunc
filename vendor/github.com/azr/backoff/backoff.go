//Package backoff helps you at backing off !
//
//It was forked from github.com/cenkalti/backoff which is awesome.
//
//This BackOff sleeps upon BackOff() and calculates its next backoff time instead of returning the duration to sleep.
package backoff

import "time"

// Interface interface to use after a retryable operation failed.
// A Interface.BackOff sleeps.
type Interface interface {
	// Example usage:
	//
	//   for ;; {
	//       err, canRetry := somethingThatCanFail()
	//       if err != nil && canRetry {
	//           backoffer.Backoff()
	//       }
	//   }
	BackOff()

	// Reset to initial state.
	Reset()
}

// ZeroBackOff is a fixed back-off policy whose back-off time is always zero,
// meaning that the operation is retried immediately without waiting.
type ZeroBackOff struct{}

var _ Interface = (*ZeroBackOff)(nil)

func (b *ZeroBackOff) Reset() {}

func (b *ZeroBackOff) BackOff() {}

type ConstantBackOff struct {
	Interval time.Duration
}

var _ Interface = (*ConstantBackOff)(nil)

func (b *ConstantBackOff) Reset() {}

func (b *ConstantBackOff) BackOff() {
	time.Sleep(b.Interval)
}

func NewConstant(d time.Duration) *ConstantBackOff {
	return &ConstantBackOff{Interval: d}
}
