package backoff

// LinearBackOff is a back-off policy whose back-off time is multiplied by mult and incremented by incr
// each time it is called.
// mult can be one ;).
import "time"

// grows linearly until
type LinearBackOff struct {
	InitialInterval time.Duration
	Multiplier      float64
	Increment       time.Duration
	MaxInterval     time.Duration
	currentInterval time.Duration
}

var _ Interface = (*LinearBackOff)(nil)

func NewLinear(from, to, incr time.Duration, mult float64) *LinearBackOff {
	return &LinearBackOff{
		InitialInterval: from,
		MaxInterval:     to,
		currentInterval: from,
		Increment:       incr,
		Multiplier:      mult,
	}
}

func (lb *LinearBackOff) Reset() {
	lb.currentInterval = lb.InitialInterval
}

func (lb *LinearBackOff) increment() {
	lb.currentInterval = time.Duration(float64(lb.currentInterval) * lb.Multiplier)
	lb.currentInterval += lb.Increment
	if lb.currentInterval > lb.MaxInterval {
		lb.currentInterval = lb.MaxInterval
	}
}

func (lb *LinearBackOff) BackOff() {
	time.Sleep(lb.currentInterval)
	lb.increment()
}
