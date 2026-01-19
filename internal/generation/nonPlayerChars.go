package generation

import "github.com/NikitaCherkashin4/go-text-game/internal/models"

func NewNPC(name string) *models.NPC {
	return &models.NPC{
		Name:   name,
		Health: 10,
	}
}
