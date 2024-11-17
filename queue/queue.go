package queue

import (
	"sync"
	"sync/atomic"
)

type Node[T any] struct {
	Value T
	next  atomic.Pointer[Node[T]]
}

type Queue[T any] struct {
	head atomic.Pointer[Node[T]]
	tail atomic.Pointer[Node[T]]
	pool sync.Pool
	len  uint32
}

func New[T any]() *Queue[T] {
	q := Queue[T]{}
	q.pool = sync.Pool{
		New: func() any {
			return &Node[T]{}
		},
	}

	return &q
}

func (q *Queue[T]) Enqueue(v T) {
	// Get이 항상 유효하기 위해 New func를 정의해야 함
	node := q.pool.Get().(*Node[T])
	node.next.Store(nil)
	node.Value = v
	headNode, tailNode := q.head.Load(), q.tail.Load()

	if tailNode != nil {
		// for{CAS} 하지 않는 이유는 어차피 현재 구현이 actor model이라 enqueue는 동시성 이슈가 없기 때문
		tailNode.next.Swap(node)
		q.tail.Swap(node)
	} else if headNode == nil {
		q.head.Store(node)
	} else {
		q.head.Load().next.Store(node)
		q.tail.Store(node)
	}

	atomic.AddUint32(&q.len, 1)
}

func (q *Queue[T]) Dequeue() *Node[T] {
	// spin형식이라 lock-contention이 높으면 안 좋을 것 같다.
	// 다만, 내 서버의 경우에는 절대 높지 않다. 고루틴 하나가
	for {
		t := q.head.Load()

		// nil을 리턴하면 아무 일도 일어나지 않을 뿐더러 pool에 nil을 넣는 쓸데없는 연산을 안하기 위해서
		// 다만, nil을 Put한다고 해서 nil이 Get 되지는 않는다. 그냥 New로 나온다.
		if t == nil {
			return t
		}

		// https://github.com/golang-design/lockfree/blob/master/queue.go
		// 하단 검사 있는 상태와 없는 상태 둘 다 벤치 돌려봤는데 큰 차이가 없다.
		if t != q.head.Load() {
			continue
		}

		if q.head.CompareAndSwap(t, t.next.Load()) {
			q.pool.Put(t)
			atomic.AddUint32(&q.len, ^uint32(0))
			return t
		}
	}
}

func (q *Queue[T]) Len() uint32 {
	return atomic.LoadUint32(&q.len)
}
