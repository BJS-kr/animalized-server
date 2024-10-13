package queue_test

import (
	"animalized/queue"
	"math/rand"
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

/*
!! WARNING !! 꼭 go test -race로 실행해주세요. 그냥 실행하면 의미가 없습니다.
*/
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

			for i := 0; i < 100; i++ {
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

type ComparisonQueue struct {
	mtx   sync.Mutex
	queue []any
}

func (q *ComparisonQueue) enqueue(v any) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.queue = append(q.queue, v)
}

func (q *ComparisonQueue) dequeue() any {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	var v any

	if len(q.queue) > 0 {
		v = q.queue[0]
		q.queue = q.queue[1:]
	}

	return v
}

// lock contention이 없는 경우, naive한 뮤텍스가 더 빠름
func BenchmarkNoContention(b *testing.B) {
	q := queue.New[any]()
	cq := ComparisonQueue{}
	n := rand.Int()

	b.ResetTimer()

	b.Run("lock free enqueue", func(b *testing.B) {
		for range b.N {
			q.Enqueue(n)
		}
	})

	b.Run("comparison enqueue", func(b *testing.B) {
		for range b.N {
			cq.enqueue(n)
		}
	})

	b.Run("lock free dequeue", func(b *testing.B) {
		for range b.N {
			q.Dequeue()
		}
	})

	b.Run("comparison dequeue", func(b *testing.B) {
		for range b.N {
			cq.dequeue()
		}
	})
}

// 구현 사항에 맞춰서 테스트: 구현 사항이란 인풋 큐에 1 actor와 1 dispatcher가 접근하는 구조
func BenchmarkContention(b *testing.B) {
	q := queue.New[any]()
	cq := ComparisonQueue{}
	n := rand.Int()

	b.ResetTimer()

	b.Run("contention lock free enqueue", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q.Enqueue(n)
			}
		})
	})

	b.Run("contention comparison enqueue", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				cq.enqueue(n)
			}
		})
	})

	b.Run("contention lock free dequeue", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q.Dequeue()
			}
		})
	})

	b.Run("contention comparison dequeue", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				cq.dequeue()
			}
		})
	})
}
