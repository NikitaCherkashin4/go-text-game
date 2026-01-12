package game

import "math/rand"

func RollD20() int {
	return rand.Intn(20) + 1
}

func RollDice(sides int) int {
	return rand.Intn(sides) + 1
}

func RollMultipleDice(numDice, sides int) int {
	total := 0
	for range numDice {
		total += rand.Intn(sides) + 1
	}
	return total
}
