package game

import "math/rand/v2"

func RollD20() int {
	return rand.IntN(20) + 1
}

func RollDice(sides int) int {
	return rand.IntN(sides) + 1
}

func RollMultipleDice(numDice, sides int) int {
	total := 0
	for range numDice {
		total += rand.IntN(sides) + 1
	}
	return total
}
