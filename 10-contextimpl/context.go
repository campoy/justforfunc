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
	"errors"
	"reflect"
	"sync"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type emptyCtx int

func (emptyCtx) Deadline() (deadline time.Time, ok bool) { return }
func (emptyCtx) Done() <-chan struct{}                   { return nil }
func (emptyCtx) Err() error                              { return nil }
func (emptyCtx) Value(key interface{}) interface{}       { return nil }

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context { return background }
func TODO() Context       { return todo }

type cancelCtx struct {
	Context
	done chan struct{}
	err  error
	mu   sync.Mutex
}

func (ctx *cancelCtx) Done() <-chan struct{} { return ctx.done }
func (ctx *cancelCtx) Err() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.err
}

var Canceled = errors.New("context canceled")

type CancelFunc func()

func WithCancel(parent Context) (Context, CancelFunc) {
	ctx := &cancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}

	cancel := func() { ctx.cancel(Canceled) }

	go func() {
		select {
		case <-parent.Done():
			ctx.cancel(parent.Err())
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}

func (ctx *cancelCtx) cancel(err error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if ctx.err != nil {
		return
	}
	ctx.err = err
	close(ctx.done)
}

type deadlineCtx struct {
	*cancelCtx
	deadline time.Time
}

func (ctx *deadlineCtx) Deadline() (deadline time.Time, ok bool) {
	return ctx.deadline, true
}

var DeadlineExceeded = deadlineExceededErr{}

type deadlineExceededErr struct{}

func (deadlineExceededErr) Error() string   { return "deadline exceeded " }
func (deadlineExceededErr) Timeout() bool   { return true }
func (deadlineExceededErr) Temporary() bool { return true }

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	cctx, cancel := WithCancel(parent)

	ctx := &deadlineCtx{
		cancelCtx: cctx.(*cancelCtx),
		deadline:  deadline,
	}

	// the line below has been replaced to work with version previous to go 1.8.
	// timeout := time.Until(deadline)
	timeout := deadline.Sub(time.Now())
	t := time.AfterFunc(timeout, func() {
		ctx.cancel(DeadlineExceeded)
	})

	stop := func() {
		t.Stop()
		cancel()
	}

	return ctx, stop
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

type valueCtx struct {
	Context
	value, key interface{}
}

func (ctx *valueCtx) Value(key interface{}) interface{} {
	if key == ctx.key {
		return ctx.value
	}
	return ctx.Context.Value(key)
}

func WithValue(parent Context, key, value interface{}) Context {
	if key == nil {
		panic("key is nil")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{
		Context: parent,
		key:     key,
		value:   value,
	}
}
