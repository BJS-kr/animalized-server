package lobby

import "animalized/handler"

func (l *Lobby) Propagate() {
	go handler.Propagate(l.inputs, l.users)
}
