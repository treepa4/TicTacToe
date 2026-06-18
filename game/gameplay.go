package game

import (
	"fmt"
	"math/rand"
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
	i := 1
	fmt.Println("  1   2   3 ")
	for row := 0; row < 3; row++ {
		if row < 2 {
			fmt.Printf("%d %s | %s | %s\n ---|---|---\n", i, g.Board[row][0], g.Board[row][1], g.Board[row][2])
			i++
		} else {
			fmt.Printf("3 %s | %s | %s\n", g.Board[row][0], g.Board[row][1], g.Board[row][2])
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
