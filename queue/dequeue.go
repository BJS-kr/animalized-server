package queue

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
		if t != q.head.Load() {
			continue
		}

		if q.head.CompareAndSwap(t, t.next.Load()) {
			q.pool.Put(t)
			return t
		}
	}
}
