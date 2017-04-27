// Copyright 2016 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package contextimpl

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestBackgroundNotTODO(t *testing.T) {
	todo := fmt.Sprint(TODO())
	bg := fmt.Sprint(Background())

	if todo == bg {
		t.Errorf("TODO and Background are equal: %q vs %q", todo, bg)
	}
}

func TestWithCancel(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil first, got %v", err)
	}
	cancel()

	<-ctx.Done()
	if err := ctx.Err(); err != Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}

func TestWithCancelConcurrent(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	time.AfterFunc(1*time.Second, cancel)

	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil first, got %v", err)
	}
	<-ctx.Done()
	if err := ctx.Err(); err != Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}

func TestWithCancelPropagation(t *testing.T) {
	ctxA, cancelA := WithCancel(Background())
	ctxB, cancelB := WithCancel(ctxA)
	defer cancelB()

	cancelA()

	select {
	case <-ctxB.Done():
	case <-time.After(1 * time.Second):
		t.Errorf("time out")
	}

	if err := ctxB.Err(); err != Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}

}

func TestWithDeadline(t *testing.T) {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := WithDeadline(Background(), deadline)

	if d, ok := ctx.Deadline(); !ok || d != deadline {
		t.Errorf("expected deadline %v; got %v", deadline, d)
	}

	then := time.Now()
	<-ctx.Done()
	if d := time.Since(then); math.Abs(d.Seconds()-2.0) > 0.1 {
		t.Errorf("should have been done after 2.0 seconds, took %v", d)
	}
	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("error should be DeadlineExceeded, got %v", err)
	}

	cancel()
	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("error should still be DeadlineExceeded, got %v", err)
	}
}

func TestWithTimeout(t *testing.T) {
	timeout := 2 * time.Second
	deadline := time.Now().Add(timeout)
	ctx, cancel := WithTimeout(Background(), timeout)

	if d, ok := ctx.Deadline(); !ok || d.Sub(deadline) > time.Millisecond {
		t.Errorf("expected deadline %v; got %v", deadline, d)
	}

	then := time.Now()
	<-ctx.Done()
	if d := time.Since(then); math.Abs(d.Seconds()-2.0) > 0.1 {
		t.Errorf("should have been done after 2.0 seconds, took %v", d)
	}
	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("error should be DeadlineExceeded, got %v", err)
	}

	cancel()
	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("error should still be DeadlineExceeded, got %v", err)
	}
}

func TestWithValue(t *testing.T) {
	tc := []struct {
		key, val, keyRet, valRet interface{}
		shouldPanic              bool
	}{
		{"a", "b", "a", "b", false},
		{"a", "b", "c", nil, false},
		{42, true, 42, true, false},
		{42, true, int64(42), nil, false},
		{nil, true, nil, nil, true},
		{[]int{1, 2, 3}, true, []int{1, 2, 3}, nil, true},
	}

	for _, tt := range tc {
		var panicked interface{}
		func() {
			defer func() { panicked = recover() }()

			ctx := WithValue(Background(), tt.key, tt.val)
			if val := ctx.Value(tt.keyRet); val != tt.valRet {
				t.Errorf("expected value %v, got %v", tt.valRet, val)
			}
		}()

		if panicked != nil && !tt.shouldPanic {
			t.Errorf("unexpected panic: %v", panicked)
		}
		if panicked == nil && tt.shouldPanic {
			t.Errorf("expected panic, but didn't get it")
		}
	}
}

func TestDeadlineExceededIsTimeouter(t *testing.T) {
	f := func(ctx Context) error {
		<-ctx.Done()
		return ctx.Err()
	}

	type timeouter interface {
		Timeout() bool
	}

	ctx, cancel := WithCancel(Background())
	cancel()
	err := f(ctx)
	if _, ok := err.(timeouter); ok {
		t.Errorf("canceled context should not have Timeout method")
	}

	ctx, cancel = WithTimeout(Background(), 1*time.Millisecond)
	defer cancel()
	err = f(ctx)
	if _, ok := err.(timeouter); !ok {
		t.Errorf("deadline exceeded context should have Timeout method")
	}
}

var (
	_ Context
	_ testing.T
)
