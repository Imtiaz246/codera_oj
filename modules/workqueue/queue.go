package workqueue

import (
	"sync"
)

type Interface[T any] interface {
	Add(item T)
	Len() int
	Next() (item T, shutdown bool)
	Done(item T)
	ShutDown()
	ShutDownWithDrain()
	ShuttingDown() bool
}

func NewWorkQueue[T any]() *Type[T] {
	return &Type[T]{
		queue:      make([]T, 0),
		dirty:      set{},
		processing: set{},
		cond:       sync.NewCond(&sync.Mutex{}),
	}
}

type Type[T any] struct {
	// queue defines the order in which we will work on items. Every
	// element of the queue should be in the dirty set and not in the
	// processing set.
	queue []T

	// dirty defines all the items that need to be processed.
	dirty set

	// Things that are currently being processed are in the processing set.
	// These things may be simultaneously in the dirty set. When we finish
	// processing something and removing it from this set, we'll check if
	// it's in the dirty set, and if so, add it to the queue.
	processing set

	cond *sync.Cond

	shuttingDown bool
	drain        bool
	nil          T
}

type empty struct{}
type t interface{}
type set map[t]empty

func (s set) has(item t) bool {
	_, exists := s[item]
	return exists
}

func (s set) insert(item t) {
	s[item] = empty{}
}

func (s set) delete(item t) {
	delete(s, item)
}

func (s set) len() int {
	return len(s)
}

// Add mark item as needing processing.
func (q *Type[T]) Add(item T) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if q.shuttingDown || q.dirty.has(item) {
		return
	}

	q.dirty.insert(item)
	if q.processing.has(item) {
		return
	}

	q.queue = append(q.queue, item)
	q.cond.Signal()
}

// Len returns the current queue length, for informational purposes only. You
// shouldn't e.g. gate a call to Add() or Get() on Len() being a particular
// value, that can't be synchronized properly.
func (q *Type[T]) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return len(q.queue)
}

// Next blocks until it can return an item to be processed. If shutdown = true,
// the caller should end their goroutine. You must call Done with item when you
// have finished processing it.
func (q *Type[T]) Next() (item T, shutdown bool) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.queue) == 0 && !q.shuttingDown {
		q.cond.Wait()
	}
	if len(q.queue) == 0 {
		// We must be shutting down.
		return q.nil, true
	}

	item = q.queue[0]
	// The underlying array still exists and references this object, so the object will not be garbage collected.
	q.queue[0] = q.nil
	q.queue = q.queue[1:]

	q.processing.insert(item)
	q.dirty.delete(item)

	return
}

// Done marks item as done processing, and if it has been marked as dirty again
// while it was being processed, it will be re-added to the queue for
// re-processing.
func (q *Type[T]) Done(item T) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.processing.delete(item)
	if q.dirty.has(item) {
		q.queue = append(q.queue, item)
		q.cond.Signal()
	} else if q.processing.len() == 0 {
		q.cond.Signal()
	}
}

// ShutDown will cause q to ignore all new items added to it and
// immediately instruct the worker goroutines to exit.
func (q *Type[T]) ShutDown() {
	q.setDrain(false)
	q.shutdown()
}

// ShutDownWithDrain will cause q to ignore all new items added to it. As soon
// as the worker goroutines have "drained," i.e.: finished processing and called
// Done on all existing items in the queue; they will be instructed to exit and
// ShutDownWithDrain will return. Hence: a strict requirement for using this is;
// your workers must ensure that Done is called on all items in the queue once
// the shut-down has been initiated, if that is not the case: this will block
// indefinitely. It is, however, safe to call ShutDown after having called
// ShutDownWithDrain, as to force the queue to shut down to terminate immediately
// without waiting for the drainage.
func (q *Type[T]) ShutDownWithDrain() {
	q.setDrain(true)
	q.shutdown()
	for q.isProcessing() && q.shouldDrain() {
		q.waitForProcessing()
	}
}

// isProcessing indicates if there are still items on the work queue being
// processed. It's used to drain the work queue on an eventual shutdown.
func (q *Type[T]) isProcessing() bool {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return q.processing.len() != 0
}

// waitForProcessing waits for the worker goroutines to finish processing items
// and call Done on them.
func (q *Type[T]) waitForProcessing() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	// Ensure that we do not wait on a queue which is already empty, as that
	// could result in waiting for Done to be called on items in an empty queue
	// which has already been shut down, which will result in waiting
	// indefinitely.
	if q.processing.len() == 0 {
		return
	}
	q.cond.Wait()
}

func (q *Type[T]) setDrain(shouldDrain bool) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.drain = shouldDrain
}

func (q *Type[T]) shouldDrain() bool {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return q.drain
}

func (q *Type[T]) shutdown() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.shuttingDown = true
	q.cond.Broadcast()
}

func (q *Type[T]) ShuttingDown() bool {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	return q.shuttingDown
}
