package queue

import "sync"

func New[T any]() *Queue[T] {
	q := Queue[T]{}
	q.pool = sync.Pool{
		New: func() any {
			return &Node[T]{}
		},
	}

	return &q
}
