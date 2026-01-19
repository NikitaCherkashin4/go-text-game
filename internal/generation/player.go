package generation

import "github.com/NikitaCherkashin4/go-text-game/internal/models"

func NewPlayer() *models.Player {
	return &models.Player{
		Name:   "Hero",
		Health: 100,
	}
}
