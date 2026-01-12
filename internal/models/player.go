package models

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Player struct {
	Name   string
	Health int
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Health: 100,
	}
}

func RequestInput(reader *bufio.Reader, message string) string {
	fmt.Print(message)
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

func GetPlayerName(input io.Reader) string {
	reader := bufio.NewReader(input)

	choice := RequestInput(reader, "Do you want to choose your own name? (yes/no): ")

	if strings.ToLower(choice) == "yes" || strings.ToLower(choice) == "y" {
		name := RequestInput(reader, "Enter your name: ")
		return name
	}

	return "Hero"
}
