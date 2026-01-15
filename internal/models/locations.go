package models

type Location struct {
	Name        string
	Description string
	HasEnemy    bool
	Enemies     []*NPC
	IsHub       bool
}
