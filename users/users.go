package users

import (
	"animalized/common"
	"animalized/message"
	"errors"
	"slices"
	"sync"
)

type Users struct {
	mtx  sync.RWMutex
	Max  int
	list []*User
}

type YieldUser func(u *User) bool

func NewUsers(maxUsers int) *Users {
	us := new(Users)
	us.list = make([]*User, 0, maxUsers)
	us.Max = maxUsers

	return us
}

func (us *Users) FindUserById(userId string) (*User, error) {
	for u := range us.LockedRange() {
		if u.Id == userId {
			return u, nil
		}
	}

	return nil, errors.New("user not found")
}

func (us *Users) Join(u *User, inputProduceChannel chan<- *message.Input) error {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	if len(us.list) >= us.Max {
		return errors.New("users max capacity reached")
	}

	u.Stop = make(chan common.Signal)
	u.SetProduceChannel(inputProduceChannel)
	us.list = append(us.list, u)

	return nil
}

func (us *Users) LockedIds() []string {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	ids := make([]string, 0, len(us.list))

	for _, u := range us.list {
		ids = append(ids, u.Id)
	}

	return ids
}

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

func (us *Users) Quit(user *User) (int, error) {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	var found *User

	us.list = slices.DeleteFunc(us.list, func(u *User) bool {
		if u == user {
			found = u

			return true
		}

		return false
	})

	if found == nil {
		return len(us.list), errors.New("failed to quit. user not found")
	}

	return len(us.list), nil
}
