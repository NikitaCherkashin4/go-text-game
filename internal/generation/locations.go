package generation

import "github.com/NikitaCherkashin4/go-text-game/internal/models"

func NewHub() *models.Location {
	return &models.Location{
		Name:        "Town Hub",
		Description: "A safe place to rest and choose your next adventure.",
		IsHub:       true,
		HasEnemy:    false,
	}
}

func NewGoblinCave() *models.Location {
	return &models.Location{
		Name:        "Goblin Cave",
		Description: "A dark cave inhabited by a goblin.",
		IsHub:       false,
		HasEnemy:    true,
		Enemies:     []*models.NPC{models.NewNPC("Goblin")},
	}
}
