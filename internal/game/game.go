package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NikitaCherkashin4/go-text-game/internal/models"
)

type Game struct {
	reader *bufio.Reader
	player *models.Player
}

func NewGame() *Game {
	fmt.Println("Welcome to the Go dungeon, Adventurer!")

	playerName := models.GetPlayerName(os.Stdin)
	player := models.NewPlayer(playerName)

	fmt.Println(*player)

	return &Game{
		reader: bufio.NewReader(os.Stdin),
		player: player,
	}
}

func (g *Game) Start() {
	fmt.Println("\nA goblin appears!")
	goblin := models.NewGoblin()

	for g.player.Health > 0 && goblin.Health > 0 {
		fmt.Printf("\n--- Your Health: %d | Goblin Health: %d ---\n",
			g.player.Health, goblin.Health)
		fmt.Println("What do you want to do?")
		fmt.Println("1. Attack")
		fmt.Println("2. Defend")
		fmt.Print("> ")

		var choice string
		fmt.Fscanln(g.reader, &choice)

		playerDamage := 0
		if choice == "1" {
			roll := RollD20()
			fmt.Printf("\nYou rolled %d!\n", roll)
			if roll >= 10 {
				playerDamage = RollDice(6)
				goblin.Health -= playerDamage
				fmt.Printf("You hit the goblin for %d damage!\n", playerDamage)
			} else {
				fmt.Println("You missed!")
			}
		} else if choice == "2" {
			fmt.Println("\nYou raise your shield...")
		}

		if goblin.Health <= 0 {
			break
		}

		enemyRoll := RollD20()
		if choice == "2" && enemyRoll < 15 {
			fmt.Printf("The goblin attacked (rolled %d) but you blocked it!\n", enemyRoll)
		} else if enemyRoll >= 8 {
			enemyDamage := RollDice(4)
			g.player.Health -= enemyDamage
			fmt.Printf("The goblin hits you for %d damage!\n", enemyDamage)
		} else {
			fmt.Printf("The goblin attacked (rolled %d) but missed!\n", enemyRoll)
		}

		fmt.Println("\nPress Enter to continue...")
		g.reader.ReadString('\n')
	}

	fmt.Println("\n=== BATTLE OVER ===")
	if g.player.Health <= 0 && goblin.Health <= 0 {
		fmt.Println("You both fell! It's a draw!")
	} else if g.player.Health <= 0 {
		fmt.Println("You were defeated! Game Over.")
	} else {
		fmt.Println("Victory! The goblin is defeated!")
		fmt.Printf("You survived with %d health remaining.\n", g.player.Health)
	}
}
