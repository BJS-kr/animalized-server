package users

func (du *DistributableUsers) Distribute() {
	for {
		select {
		case <-du.Stop:
			return
		default:
			n := du.Inputs.Dequeue()

			if n == nil {
				continue
			}

			for u := range du.Users.LockedRange() {
				u.Inputs.Enqueue(n.Value)
			}
		}
	}
}
