package message

import "sync"

type MsgPool struct {
	pool sync.Pool
}

var Pool = MsgPool{
	pool: sync.Pool{
		New: func() any {
			return new(Input)
		},
	},
}

func (mp *MsgPool) Get() *Input {
	i := mp.pool.Get().(*Input)
	i.Reset()

	return i
}

func (mp *MsgPool) Put(i *Input) {
	mp.pool.Put(i)
}
