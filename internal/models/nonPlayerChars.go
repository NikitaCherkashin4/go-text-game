package models

type NPC struct {
	Name   string
	Health int
}

func NewNPC(Name string) *NPC {
	return &NPC{Health: 10}
}
