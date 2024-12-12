package users

import (
	"animalized/message"
	"errors"
	"slices"
	"sync"
)

type Session struct {
	mtx  sync.RWMutex
	Max  int
	list []*User
}

type YieldUser func(u *User) bool

func NewSession(maxUsers int) *Session {
	ss := new(Session)
	ss.list = make([]*User, 0, maxUsers)
	ss.Max = maxUsers

	return ss
}

func (ss *Session) FindUserById(userId string) (*User, error) {
	for u := range ss.LockedRange() {
		if u.Id == userId {
			return u, nil
		}
	}

	return nil, errors.New("user not found")
}

func (ss *Session) Join(u *User, inputProduceChannel chan<- *message.Input) error {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	if len(ss.list) >= ss.Max {
		return errors.New("users max capacity reached")
	}

	u.SetProduceChannel(inputProduceChannel)
	ss.list = append(ss.list, u)

	return nil
}

func (ss *Session) LockedIds() []string {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	ids := make([]string, 0, len(ss.list))

	for _, u := range ss.list {
		ids = append(ids, u.Id)
	}

	return ids
}

func (ss *Session) LockedRange() func(YieldUser) {
	return func(yield YieldUser) {
		ss.mtx.RLock()
		defer ss.mtx.RUnlock()

		for _, u := range ss.list {
			if !yield(u) {
				return
			}
		}
	}
}

func (ss *Session) Quit(user *User) (int, error) {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	var found *User

	ss.list = slices.DeleteFunc(ss.list, func(u *User) bool {
		if u == user {
			found = u

			return true
		}

		return false
	})

	if found == nil {
		return len(ss.list), errors.New("failed to quit. user not found")
	}

	return len(ss.list), nil
}
