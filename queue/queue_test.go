package queue_test

import (
	"animalized/queue"
	"sync"
	"testing"
)

func repeatEnqueue(q *queue.Queue[int], repeat int) {
	for i := 0; i < repeat; i++ {
		q.Enqueue(i)
	}
}

func dequeueToDrain(q *queue.Queue[int]) {
	for {
		if q.Dequeue() == nil {
			break
		}
	}
}

func TestBasicOperation(t *testing.T) {
	tcs := []struct {
		input int
	}{
		{
			input: 1,
		},
		{
			input: 100_000,
		},
	}

	for _, tc := range tcs {
		q := queue.New[int]()
		repeatEnqueue(q, tc.input)

		var v int
		for {
			n := q.Dequeue()

			if n != nil {
				if n.Value < v {
					t.Fatal("node didn't pulled out orderly")
				}
				v = n.Value
				continue
			}

			break
		}
	}
}

func TestOperationAfterDrain(t *testing.T) {
	q := queue.New[int]()

	tcs := []struct {
		input int
	}{
		{
			input: 1,
		},
		{
			input: 100_000,
		},
	}

	for _, tc := range tcs {
		repeatEnqueue(q, tc.input)
		dequeueToDrain(q)
		repeatEnqueue(q, tc.input)
		dequeueToDrain(q)
	}
}

func TestDequeueBeforeEnqueue(t *testing.T) {
	q := queue.New[int]()
	if q.Dequeue() != nil {
		t.Fatal("dequeue before enqueue failed")
	}
}

// go test -race로 실행해주세요
func TestRaceCondition(t *testing.T) {
	var outerWg sync.WaitGroup
	input := 1_000_000

	for r := 0; r < 10; r++ {
		outerWg.Add(1)
		go func() {
			var wg sync.WaitGroup
			q := queue.New[int]()

			wg.Add(1)
			// ! 현재 서버 구현상 insert가 무조건 single thread로 일어나기 때문에 insert는 goroutine 하나
			go func() {
				repeatEnqueue(q, input)
				wg.Done()
			}()

			for i := 0; i < 10; i++ {
				wg.Add(1)
				go func() {
					for {
						if q.Dequeue() == nil {
							break
						}
					}
					wg.Done()
				}()
			}

			wg.Wait()
			outerWg.Done()
		}()
	}
	outerWg.Wait()
}
