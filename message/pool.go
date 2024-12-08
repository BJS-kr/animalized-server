package message

import "sync"

type MsgPool struct {
	pool sync.Pool
}

func (mp *MsgPool) Get() *Input {
	return mp.pool.Get().(*Input)
}

func (mp *MsgPool) Put(i *Input) {
	i.Reset()
	mp.pool.Put(i)
}

var Pool = &MsgPool{
	pool: sync.Pool{
		New: func() any {
			return new(Input)
		},
	},
}
