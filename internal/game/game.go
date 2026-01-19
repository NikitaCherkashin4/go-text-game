package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NikitaCherkashin4/go-text-game/internal/generation"
	"github.com/NikitaCherkashin4/go-text-game/internal/models"
)

func NewGame() *models.Game {
	fmt.Println("Welcome to the Go dungeon, Adventurer!")

	reader := bufio.NewReader(os.Stdin)

	player := generation.NewPlayer()
	player.SetName(reader)

	hub := generation.NewHub()
	goblinCave := generation.NewGoblinCave()

	gameState := &models.GameState{
		CurrentLocation:    hub,
		Hub:                hub,
		AvailableLocations: []*models.Location{goblinCave},
	}

	return &models.Game{
		Reader: reader,
		Player: player,
		State:  gameState,
	}
}
