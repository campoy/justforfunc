package anaconda

import (
	"time"

	"github.com/azr/backoff"
)

/*
Reconnecting(from https://dev.twitter.com/streaming/overview/connecting) :

Once an established connection drops, attempt to reconnect immediately.
If the reconnect fails, slow down your reconnect attempts according to the type of error experienced:
*/

//Back off linearly for TCP/IP level network errors.
//	These problems are generally temporary and tend to clear quickly.
//	Increase the delay in reconnects by 250ms each attempt, up to 16 seconds.
func NewTCPIPErrBackoff() backoff.Interface {
	return backoff.NewLinear(0, time.Second*16, time.Millisecond*250, 1)
}

//Back off exponentially for HTTP errors for which reconnecting would be appropriate.
//	Start with a 5 second wait, doubling each attempt, up to 320 seconds.
func NewHTTPErrBackoff() backoff.Interface {
	eb := backoff.NewExponential()
	eb.InitialInterval = time.Second * 5
	eb.MaxInterval = time.Second * 320
	eb.Multiplier = 2
	eb.Reset()
	return eb
}

// Back off exponentially for HTTP 420 errors.
// 	Start with a 1 minute wait and double each attempt.
// 	Note that every HTTP 420 received increases the time you must
// 	wait until rate limiting will no longer will be in effect for your account.
func NewHTTP420ErrBackoff() backoff.Interface {
	eb := backoff.NewExponential()
	eb.InitialInterval = time.Minute * 1
	eb.Multiplier = 2
	eb.MaxInterval = time.Minute * 20
	eb.Reset()
	return eb
}
