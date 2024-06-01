package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var name string
	for {
		fmt.Println("What is your first name?")
		numberInput, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if numberInput == 1 {
			break
		}
	}

	computer := newPlayer("computer")
	player1 := newPlayer(name)

	var decisionOnBoard string
	computer.board = newBoard()
	game := newGame()
	addPlayerToGame(game, computer)
	for {
		player1.board = newBoard()
		fmt.Println(player1.board)
		fmt.Println("Are you happy with this board y/n")
		fmt.Println("If you want a new board say n, if you're happy with it say n")
		numberInput, err := fmt.Scanln(&decisionOnBoard)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if numberInput == 1 && decisionOnBoard == "y" {
			break
		}
	}

	fmt.Println(player1)
	fmt.Println(computer)

	addPlayerToGame(game, player1)

	for !player1.winner && !computer.winner {
		Letter, Column, Number := caller()
		fmt.Println(Letter, Column, Number)
		for i := 0; i < 5; i++ {
			if player1.board[i][Column] == strconv.Itoa(Number) {
				player1.board[i][Column] = "x"
			}
			if computer.board[i][Column] == strconv.Itoa(Number) {
				computer.board[i][Column] = "x"
			}
		}
		player1.checkWon()
		computer.CheckWon()

	}
}
