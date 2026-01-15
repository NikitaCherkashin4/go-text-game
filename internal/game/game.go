package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NikitaCherkashin4/go-text-game/internal/generation"
	"github.com/NikitaCherkashin4/go-text-game/internal/models"
)

type Game struct {
	reader *bufio.Reader
	player *models.Player
	state  *GameState
}

type GameState struct {
	CurrentLocation    *models.Location
	Hub                *models.Location
	AvailableLocations []*models.Location
}

func NewGame() *Game {
	fmt.Println("Welcome to the Go dungeon, Adventurer!")

	reader := bufio.NewReader(os.Stdin)

	player := models.NewPlayer()
	player.SetName(reader)

	hub := generation.NewHub()
	goblinCave := generation.NewGoblinCave()

	gameState := &GameState{
		CurrentLocation:    hub,
		Hub:                hub,
		AvailableLocations: []*models.Location{goblinCave},
	}

	for i, enemy := range goblinCave.Enemies {
		fmt.Println(i, enemy)
	}

	return &Game{
		reader: reader,
		player: player,
		state:  gameState,
	}
}

func (g *Game) Start() {
	fmt.Println(g.state.AvailableLocations)
}
