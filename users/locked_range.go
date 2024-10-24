package users

// consumer.Propagate에서 Users.users range를 돌아야 한다.
// 다른 고루틴이 Users.users를 간섭하므로 lock이 필요하다.
// lock-free로 구현하지 않은 이유는 경합이 낮지만 단순히 정합성을 유지하기 위해 락을 걸기 때문이다.
// lock을 즉시 획득할 수 있는 경우가 대부분이라면 lock-free보다 mutex가 빠르고 간편하다.

type YieldUser func(u *User) bool

func (us *Users) LockedRange() func(YieldUser) {
	return func(yield YieldUser) {
		us.mtx.RLock()
		defer us.mtx.RUnlock()

		for _, u := range us.list {
			if !yield(u) {
				return
			}
		}
	}
}
