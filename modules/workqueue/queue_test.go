package workqueue

import (
	"sync"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	tests := []struct {
		queue         *Type[int]
		queueShutDown func(Interface[int])
	}{
		{
			queue:         NewWorkQueue[int](),
			queueShutDown: Interface[int].ShutDown,
		},
		{
			queue:         NewWorkQueue[int](),
			queueShutDown: Interface[int].ShutDownWithDrain,
		},
	}

	for _, test := range tests {
		// If something is seriously wrong, this function will never be completed.

		// Start producers
		const producers = 50
		producerWG := sync.WaitGroup{}
		producerWG.Add(producers)

		for i := 1; i <= producers; i++ {
			go func(i int) {
				defer producerWG.Done()
				for j := 1; j <= 50; j++ {
					test.queue.Add(j)
					time.Sleep(time.Millisecond)
				}
			}(i)
		}

		// Start consumers
		const consumers = 10
		consumerWG := sync.WaitGroup{}
		consumerWG.Add(consumers)

		for i := 1; i <= consumers; i++ {
			go func(i int) {
				defer consumerWG.Done()
				for {
					item, quit := test.queue.Next()
					if item == -1 {
						t.Errorf("Got an item added after shutdown!")
					}

					if quit {
						return
					}

					t.Logf("Worker %v: begin processing %v", i, item)
					time.Sleep(3 * time.Millisecond)
					t.Logf("Worker %v: done processing %v", i, item)
					test.queue.Done(item)
				}
			}(i)
		}

		producerWG.Wait()
		test.queueShutDown(test.queue)
		test.queue.Add(-1) // added after shutdown
		consumerWG.Wait()

		if test.queue.Len() != 0 {
			t.Errorf("Expected the queue to be empty, has %v items", test.queue.Len())
		}
	}
}
