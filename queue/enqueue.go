package queue

import "sync/atomic"

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
