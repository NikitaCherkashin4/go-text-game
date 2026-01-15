package game

import (
	"bufio"
	"fmt"
	"os"

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

	hub := &models.Location{
		Name:        "Town Hub",
		Description: "A safe place to rest and choose your next adventure.",
		IsHub:       true,
		HasEnemy:    false,
	}

	goblinCave := &models.Location{
		Name:        "Goblin Cave",
		Description: "A dark cave inhabited by a goblin.",
		IsHub:       false,
		HasEnemy:    true,
		Enemies:     []*models.NPC{models.NewNPC("Goblin")},
	}

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

func (g *Game) Start() {}
