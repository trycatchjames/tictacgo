//+build !test

package main

import (
	"bufio"
	"fmt"
	"os"
)

func instructions() {
	fmt.Println("")
	fmt.Println("Enter a number between 1 and 9:")
	fmt.Println("1 2 3")
	fmt.Println("4 5 6")
	fmt.Println("7 8 9")
	fmt.Println("")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	isGameFinished := false
	winner := ""
	g := new(Game)

	fmt.Println("Welcome to tic tac to")
	for !isGameFinished {
		instructions()
		fmt.Println(g)
		scanner.Scan()
		selection := scanner.Text()
		if !g.Move(selection) {
			fmt.Println("Invalid input! Enter a number between 1 - 9")
			fmt.Println("1 is top left and 9 is bottom right")
			continue
		}
		isGameFinished, winner = g.CheckResults()

		if !isGameFinished {
			g.ComputerMove()
			isGameFinished, winner = g.CheckResults()
		}
	}

	fmt.Println(g)

	if winner == "" {
		fmt.Println("Draw!")
	} else if winner == "x" {
		fmt.Println("You win :D")
	} else {
		fmt.Println("You lose :(")
	}
}
