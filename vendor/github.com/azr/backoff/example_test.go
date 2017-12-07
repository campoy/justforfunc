package backoff_test

import (
	"fmt"
	"time"

	"github.com/azr/backoff"
)

func ExampleNewExponential_defaultWaitingIntervals() {
	exp := backoff.NewExponential()

	for i := 0; i < 25; i++ {
		d := exp.GetSleepTime()
		fmt.Printf("Random duration was %2.2fs, interval: %2.2fs in [ %2.2fs , %2.2fs ]\n",
			d.Seconds(),
			exp.Inverval().Seconds(),
			(exp.Inverval() - time.Duration(exp.RandomizationFactor*float64(exp.Inverval()))).Seconds(),
			(exp.Inverval() + time.Duration(exp.RandomizationFactor*float64(exp.Inverval()))).Seconds(),
		)
		exp.IncrementCurrentInterval()
		// exp.BackOff() would have executed time.Sleep(exp.GetSleepTime()) and exp.IncrementCurrentInterval()
	}
	// Output:
	// Random duration was 0.51s, interval: 0.50s in [ 0.25s , 0.75s ]
	// Random duration was 0.99s, interval: 0.75s in [ 0.38s , 1.12s ]
	// Random duration was 0.80s, interval: 1.12s in [ 0.56s , 1.69s ]
	// Random duration was 1.49s, interval: 1.69s in [ 0.84s , 2.53s ]
	// Random duration was 2.07s, interval: 2.53s in [ 1.27s , 3.80s ]
	// Random duration was 3.68s, interval: 3.80s in [ 1.90s , 5.70s ]
	// Random duration was 4.46s, interval: 5.70s in [ 2.85s , 8.54s ]
	// Random duration was 6.78s, interval: 8.54s in [ 4.27s , 12.81s ]
	// Random duration was 15.11s, interval: 12.81s in [ 6.41s , 19.22s ]
	// Random duration was 13.81s, interval: 19.22s in [ 9.61s , 28.83s ]
	// Random duration was 20.27s, interval: 28.83s in [ 14.42s , 43.25s ]
	// Random duration was 37.23s, interval: 43.25s in [ 21.62s , 64.87s ]
	// Random duration was 64.24s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 81.75s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 47.59s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 47.82s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 75.15s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 42.39s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 81.92s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 71.80s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 61.43s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 31.70s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 39.50s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 66.44s, interval: 60.00s in [ 30.00s , 90.00s ]
	// Random duration was 88.51s, interval: 60.00s in [ 30.00s , 90.00s ]
}
