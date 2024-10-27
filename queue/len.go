package queue

import "sync/atomic"

func (q *Queue[T]) Len() uint32 {
	return atomic.LoadUint32(&q.len)
}
