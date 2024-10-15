package state

type UserState struct {
	position *Position
	score    int32
}

func (us *UserState) IncreaseUserScore(amount int32) {
	us.score += amount
}
