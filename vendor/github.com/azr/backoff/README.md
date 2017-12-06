# backoff

[![GoDoc](https://godoc.org/github.com/azr/backoff?status.png)](https://godoc.org/github.com/azr/backoff)
[![Build Status](https://travis-ci.org/azr/backoff.png)](https://travis-ci.org/azr/backoff)

This is a fork from the awesome [cenkalti/backoff](github.com/cenkalti/backoff) which is a go port from
[google-http-java-client](https://code.google.com/p/google-http-java-client/wiki/ExponentialBackoff).

This BackOff sleeps upon BackOff() and calculates its next backoff time instead of returning the duration to sleep.

[Exponential backoff](http://en.wikipedia.org/wiki/Exponential_backoff)
is an algorithm that uses feedback to multiplicatively decrease the rate of some process,
in order to gradually find an acceptable rate.
The retries exponentially increase and stop increasing when a certain threshold is met.



## Install

```bash
go get github.com/azr/backoff
```
