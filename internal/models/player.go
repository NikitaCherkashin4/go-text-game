package models

import (
	"bufio"
	"fmt"
	"strings"
)

type Player struct {
	Name   string
	Health int
}

func RequestInput(reader *bufio.Reader, message string) string {
	fmt.Print(message)
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

func (p *Player) SetName(reader *bufio.Reader) {
	choice := RequestInput(reader, "Do you want to choose your own name? (yes/no): ")

	if strings.ToLower(choice) == "yes" || strings.ToLower(choice) == "y" {
		name := RequestInput(reader, "Enter your name: ")
		p.Name = name
	}
}
