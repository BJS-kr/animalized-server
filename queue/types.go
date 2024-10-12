package queue

import (
	"sync"
	"sync/atomic"
)

// atomic.Value도 선택지에 있었지만 타이핑이 안된다(제네릭 안받고 무조건 any임).
// https://stackoverflow.com/questions/64938715/when-should-we-choose-locking-over-lock-free-data-structures
// 위 글에 따르면 lock-free는 만능이 아닐뿐더러,high-contention에서는 쓸 수 없다.
// https://stackoverflow.com/questions/1585818/when-are-lock-free-data-structures-less-performant-than-mutual-exclusion-mutexe
// 위 글이 왜 high contention에서 lock-free를 쓸 수 없는지 설명해주고 있다. sync.Pool로 해결가능?
// https://medium.com/@tylerneely/fear-and-loathing-in-lock-free-programming-7158b1cdd50c
// 유지보수성도 신경써야한다. 간단한 시나리오에서는 lock이 더 명확하다. 그러나 lock이 중첩되기 시작하면? 얘기가 또 다르다. 이건 정해진 건 없다.
// 결론적으로 내 구현체에서 ABA problem은 없다. 포인터라서 없다는 말이 아니다. 연산 자체가 monotonic하기 때문이다(멀티스레딩 전문가가 아니라서 100%인지는 모르겠음...).
// TODO test와 benchmark 작성해서 진짜 문제가 없고 써도 된다는 것 확신
type Node[T any] struct {
	value T
	next  atomic.Pointer[Node[T]]
}

type Queue[T any] struct {
	head atomic.Pointer[Node[T]]
	tail atomic.Pointer[Node[T]]
	pool sync.Pool
}
