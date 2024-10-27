package utility

type Level struct {
	Actors []*Actor
	Pawns  []*Actor
}

func NewLevel() *Level {
	return &Level{}
}
