package models

type NPC struct {
	Name   string
	Health int
}

func NewNPC(name string) *NPC {
	return &NPC{
		Name:   name,
		Health: 10,
	}
}
