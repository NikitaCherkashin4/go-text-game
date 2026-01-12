package models

type Goblin struct {
	Health int
}

func NewGoblin() *Goblin {
	return &Goblin{Health: 10}
}
