package common

func (b *Base) Propagate() {
	for {
		select {
		case <-b.Stop:
			return
		default:
			n := b.Inputs.Dequeue()

			if n == nil {
				continue
			}

			for u := range b.Users.LockedRange() {
				u.Inputs.Enqueue(n.Value)
			}
		}
	}
}
