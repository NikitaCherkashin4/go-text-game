package game

import "github.com/NikitaCherkashin4/go-text-game/internal/models"

func MoveToLocation(player *models.Player, location *models.Location) {
	// TODO
}

func ReturnToHub(player *models.Player, hub *models.Location) {
	// TODO
}

func RunLocation(player *models.Player, location *models.Location, hub *models.Location) {
	// TODO
	ReturnToHub(player, hub)
}
