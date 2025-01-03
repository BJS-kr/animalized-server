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

	var headNode *Node[T]
	var tailNode *Node[T]

	for {
		headNode = q.head.Load()
		tailNode = q.tail.Load()

		// 이 지점에서 Dequeue가 일어나 head가 바뀌었을 수 있음
		if headNode == nil && q.head.Load() == nil && q.head.CompareAndSwap(headNode, node) {

			// head가 비었으므로 비교 없이 강제로 store되어야 한다.
			q.tail.Store(node)
			break
			// tail이 바뀌지 않았고 tailNode의 next가 nil인지(진짜 tail인지) 확인
		} else if tailNode == q.tail.Load() && tailNode != nil && tailNode.next.Load() == nil {
			// Node가 한 개만 존재할 때 head와 tail이 같으므로 tailNode의 next에 삽입하는 것은 head의 next에 삽입하는 것과 같다.
			tailNode.next.Store(node)
			q.tail.Store(node)
			break
		}
	}

	atomic.AddUint32(&q.len, 1)
}

func (q *Queue[T]) Dequeue() *Node[T] {
	// spin형식이라 lock-contention이 높으면 안 좋을 것 같다.
	// 다만, 내 서버의 경우에는 절대 높지 않다. 고루틴 하나가
	for {
		n := q.head.Load()

		// nil을 리턴하면 아무 일도 일어나지 않을 뿐더러 pool에 nil을 넣는 쓸데없는 연산을 안하기 위해서
		// 다만, nil을 Put한다고 해서 nil이 Get 되지는 않는다. 그냥 New로 나온다.
		if n == nil {
			return n
		}

		// https://github.com/golang-design/lockfree/blob/master/queue.go
		// 하단 검사 있는 상태와 없는 상태 둘 다 벤치 돌려봤는데 큰 차이가 없다.
		if n != q.head.Load() {
			continue
		}

		if q.head.CompareAndSwap(n, n.next.Load()) {
			atomic.AddUint32(&q.len, ^uint32(0))

			if n == q.tail.Load() {
				q.tail.Store(nil)
			}

			return n
		}
	}
}

func (q *Queue[T]) Put(n *Node[T]) {
	q.pool.Put(n)
}

func (q *Queue[T]) GetHead() *Node[T] {
	return q.head.Load()
}

func (q *Queue[T]) GetTail() *Node[T] {
	return q.tail.Load()
}

func (q *Queue[T]) Len() uint32 {
	return atomic.LoadUint32(&q.len)
}
