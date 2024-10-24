package users

func (du *DistributableUsers) StopStreaming() {
	close(du.Stop)
}
