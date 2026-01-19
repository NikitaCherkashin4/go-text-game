package models

import (
	"bufio"
)

type Game struct {
	Reader *bufio.Reader
	Player *Player
	State  *GameState
}

func (g *Game) Start() {}

type GameState struct {
	CurrentLocation    *Location
	Hub                *Location
	AvailableLocations []*Location
}
