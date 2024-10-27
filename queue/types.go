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
