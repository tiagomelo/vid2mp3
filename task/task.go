// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package task

import (
	"context"
	"sync"
)

// Worker must be implemented by types that want to use
// the run pool.
type Worker interface {
	Work(ctx context.Context) error
}

// Task provides a pool of goroutines that can execute any Worker
// tasks that are submitted.
type Task struct {
	work    chan Worker
	wg      sync.WaitGroup
	errChan chan error
}

// New creates a new work pool.
func New(ctx context.Context, maxGoroutines, taskQueueSize int) *Task {
	t := Task{
		work:    make(chan Worker, taskQueueSize),
		errChan: make(chan error, taskQueueSize),
	}

	// The goroutines are the pool. So we could add code
	// to change the size of the pool later on.

	t.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range t.work {
				if err := w.Work(ctx); err != nil {
					t.errChan <- err
				}
			}
			t.wg.Done()
		}()
	}

	return &t
}

// Shutdown waits for all the goroutines to shutdown.
func (t *Task) Shutdown() {
	close(t.work)
	t.wg.Wait()
	close(t.errChan)
}

// Errors returns a channel that will receive errors
// that occurred while processing the tasks.
func (t *Task) Errors() <-chan error {
	return t.errChan
}

// Do submits work to the pool.
func (t *Task) Do(w Worker) {
	t.work <- w
}
