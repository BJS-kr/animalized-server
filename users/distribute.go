package users

func (b *DistributableUsers) Distribute() {
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
