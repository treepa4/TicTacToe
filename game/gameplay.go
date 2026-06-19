package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type Player struct {
	Name   string
	Symbol string
}

type Game struct {
	Board         [3][3]string
	HumanPlayer   Player
	AiPlayer      Player
	CurrentPlayer Player
	AiWinCount    int
	HumanWinCount int
}

func (g *Game) AiCall() {
	fmt.Println("Ход компьютера...")
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if g.Board[row][col] == "-" {
			g.Board[row][col] = g.AiPlayer.Symbol
			break
		}
	}
	return
}

func (g *Game) HumanCall(row, col int) bool {
	if row > 2 || col > 2 || row < 0 || col < 0 {
		fmt.Println("Неверные координаты")
		return false
	}
	if g.Board[row][col] == "-" {
		g.Board[row][col] = g.HumanPlayer.Symbol
		return true
	}
	fmt.Println("Клетка занята")
	return false
}

func NewGame(HumanSymbol string) *Game {
	aiSymbol := "o"
	if HumanSymbol == "o" {
		aiSymbol = "x"
	}
	g := &Game{
		HumanPlayer: Player{
			Name:   "Игрок",
			Symbol: HumanSymbol,
		},
		AiPlayer: Player{
			Name:   "Компьютер",
			Symbol: aiSymbol,
		},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			g.Board[i][j] = "-"
		}
	}
	if g.HumanPlayer.Symbol == "x" {
		g.CurrentPlayer = g.HumanPlayer
	} else {
		g.CurrentPlayer = g.AiPlayer
	}
	return g
}

func (g *Game) PrintBoard() {
	fmt.Println("  1   2   3 ")
	for row := 0; row < 3; row++ {
		if row < 2 {
			fmt.Printf("%d %s | %s | %s\n ---|---|---\n", row+1, g.Board[row][0], g.Board[row][1], g.Board[row][2])
		} else {
			fmt.Printf("%d %s | %s | %s\n", row+1, g.Board[row][0], g.Board[row][1], g.Board[row][2])
		}
	}
}

func (g *Game) CheckWin(s string) bool {
	for i := 0; i < 3; i++ {
		if (g.Board[i][0] == s && g.Board[i][1] == s && g.Board[i][2] == s) ||
			(g.Board[0][i] == s && g.Board[1][i] == s && g.Board[2][i] == s) {
			return true
		}
	}
	if (g.Board[0][0] == s && g.Board[1][1] == s && g.Board[2][2] == s) ||
		(g.Board[0][2] == s && g.Board[1][1] == s && g.Board[2][0] == s) {
		return true
	}
	return false
}

func (g *Game) CheckDraw() bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if g.Board[row][col] == "-" {
				return false
			}
		}
	}
	return true
}

func (g *Game) PrintScore() {
	fmt.Println("Отличная партия! Вот ваш счет:")
	fmt.Printf("Вы: %d\n", g.HumanWinCount)
	fmt.Printf("Протвник: %d\n", g.AiWinCount)
}

func StartGame(g *Game) {
	for {
		if g.CurrentPlayer.Name == "Игрок" {
			var row, col int
			for {
				fmt.Println("Твой ход! Напиши номер строки и номер столбца через пробел")
				g.PrintBoard()
				_, err := fmt.Scanln(&row, &col)
				if err != nil {
					fmt.Println("Пожалуйста, введите два числа через пробел (например: 1 2)")
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
			if g.CurrentPlayer.Name == "Игрок" {
				g.HumanWinCount++
				if AskNewGame(g) {
					g.Reset()
					continue
				} else {
					return
				}

			}
			g.AiWinCount++
			if AskNewGame(g) {
				g.Reset()
				continue
			} else {
				return
			}

		}
		if g.CheckDraw() {
			fmt.Println("Партия завершилась ничьей!")
			g.HumanWinCount++
			g.AiWinCount++
			if AskNewGame(g) {
				g.Reset()
				continue
			} else {
				return
			}

		}
		if g.CurrentPlayer.Name == "Игрок" {
			g.CurrentPlayer = g.AiPlayer
		} else {
			g.CurrentPlayer = g.HumanPlayer
		}

	}

}

func (g *Game) Reset() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			g.Board[i][j] = "-"
		}
	}
	if g.HumanPlayer.Symbol == "x" {
		g.CurrentPlayer = g.HumanPlayer
	} else {
		g.CurrentPlayer = g.AiPlayer
	}
}
func AskNewGame(g *Game) bool {
	fmt.Println("Желаете ли вы продолжить? (y/n): ")
	s := bufio.NewScanner(os.Stdin)
	if ok := s.Scan(); !ok {
		fmt.Println("Error reading input")
		return false
	}
	answer := s.Text()
	switch answer {
	case "y":
		if g.HumanPlayer.Symbol == "x" {
			g.HumanPlayer.Symbol = "o"
			g.AiPlayer.Symbol = "x"
		} else {
			g.HumanPlayer.Symbol = "x"
			g.AiPlayer.Symbol = "o"
		}
		return true
	case "n":
		g.PrintScore()
	}
	return false
}
