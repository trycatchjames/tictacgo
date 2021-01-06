package main

import (
	"math/rand"
	"strconv"
)

// Game board and state
type Game struct {
	Board [9]string
}

// Draw the board to std out
func (g *Game) String() string {
	buffer := "-------------\n|"
	for index, element := range g.Board {
		counter := index + 1
		if element == "" {
			buffer += "   |"
		} else {
			buffer += " " + element + " |"
		}
		if counter%3 == 0 && counter == 9 {
			buffer += "\n-------------"
		} else if counter%3 == 0 {
			buffer += "\n-------------\n|"
		}
	}
	return buffer
}

// Move the human
func (g *Game) Move(input string) bool {
	i, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	if i < 1 || i > 9 {
		return false
	}
	if g.Board[i-1] != "" {
		return false
	}
	g.Board[i-1] = "x"
	return true
}

func (g *Game) checkRow(ref1 int, ref2 int, ref3 int) bool {
	return g.Board[ref1] != "" && g.Board[ref1] == g.Board[ref2] && g.Board[ref1] == g.Board[ref3]
}

/*
CheckResults to see if there's a win

0, 1, 2
3, 4, 5
6, 7, 8
*/
func (g *Game) CheckResults() (bool, string) {
	if g.checkRow(0, 1, 2) {
		// top row
		return true, g.Board[0]
	}
	if g.checkRow(3, 4, 5) {
		// middle row
		return true, g.Board[3]
	}
	if g.checkRow(6, 7, 8) {
		// bottom row
		return true, g.Board[6]
	}
	if g.checkRow(0, 3, 6) {
		// left column
		return true, g.Board[0]
	}
	if g.checkRow(1, 4, 7) {
		// centre column
		return true, g.Board[1]
	}
	if g.checkRow(2, 5, 8) {
		// right column
		return true, g.Board[2]
	}
	if g.checkRow(0, 4, 8) {
		// diagonal top left
		return true, g.Board[0]
	}
	if g.checkRow(6, 4, 2) {
		// diagonal bottom left
		return true, g.Board[2]
	}
	for _, element := range g.Board {
		if element == "" {
			return false, ""
		}
	}
	return true, ""
}

// ComputerMove will make the computers next turn
func (g *Game) ComputerMove() {
	for {
		selection := rand.Intn(len(g.Board))
		if g.Board[selection] == "" {
			g.Board[selection] = "o"
			return
		}
	}
}
