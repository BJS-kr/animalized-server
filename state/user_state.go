package state

type UserState struct {
	Position *Position
	score    int32
}

func (us *UserState) IncreaseUserScore(amount int32) {
	us.score += amount
}

func (us *UserState) IsWinner() bool {
	return us.score >= WIN_SCORE
}
