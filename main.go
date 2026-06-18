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
	fmt.Println("Игровое поле: ")
	g.PrintBoard()
	for {
		if g.CurrentPlayer.Name == "Игрок" {
			var row, col int
			for {
				fmt.Println("Твой ход! Напиши номер строки и номер столбца через пробел")
				_, err := fmt.Scanln(&row, &col)
				if err != nil {
					fmt.Println("Пожалуйста, введите два числа через пробел (например: 1 2)")
					var discard string
					fmt.Scanln(&discard)
					continue
				}
				if g.HumanCall(row-1, col-1) {
					break
				}

			}
		} else {
			g.AiCall()
		}
		g.PrintBoard()
		if g.CheckWin(g.CurrentPlayer.Symbol) {
			fmt.Println("Партия завершилась победой ", g.CurrentPlayer.Name)
			break
		}
		if g.CheckDraw() {
			fmt.Println("Партия завершилась ничьей!")
			break
		}
		if g.CurrentPlayer.Name == "Игрок" {
			g.CurrentPlayer = g.AiPlayer
		} else {
			g.CurrentPlayer = g.HumanPlayer
		}
	}
}
