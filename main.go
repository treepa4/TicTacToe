package main

import (
	"TicTacToe/game"
	"TicTacToe/greeting"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	greeting.Greeting()

	fmt.Printf("Введите за кого вы будете играть(x/o): ")
	scan := bufio.NewScanner(os.Stdin)
	if ok := scan.Scan(); !ok {
		fmt.Println("Scan failed")
		return
	}

	humanS := scan.Text()
	switch humanS {
	case "х":
		humanS = "x"
	case "0":
		humanS = "o"
	case "о":
		humanS = "o"
	case "1":
		humanS = "x"
	case "2":
		humanS = "o"
	default:
		humanS = strings.TrimSpace(strings.ToLower(humanS))
	}
	g := game.NewGame(humanS)
	g.HumanWinCount = 0
	g.AiWinCount = 0
	game.StartGame(g)

}
