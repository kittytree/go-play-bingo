package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
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
	for {
		player1.board = newBoard()
		playerBoardToString(player1)
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
	fmt.Printf("%s board:\n", player1.name)
	playerBoardToString(player1)
	fmt.Println("Computer board:")
	playerBoardToString(computer)

	for !player1.winner && !computer.winner {

		Letter, Column, Number := caller(game)
		fmt.Printf("Does anyone have: %s,%v?\n", Letter, Number)
		time.Sleep(1 * time.Second)
		for i := 0; i < 5; i++ {
			if player1.board[i][Column] == strconv.Itoa(Number) {
				player1.board[i][Column] = " X"
				fmt.Printf("Match found! for %s\n", player1.name)
				playerBoardToString(player1)
				time.Sleep(1 * time.Second)

			}
			if computer.board[i][Column] == strconv.Itoa(Number) {
				computer.board[i][Column] = " X"
				fmt.Println("Match found! for Computer")
				playerBoardToString(computer)
				time.Sleep(1 * time.Second)
			}
		}
		if checkWon(player1) {
			player1.won()
		}
		if checkWon(computer) {
			computer.won()
		}
	}
	if player1.winner && computer.winner {
		fmt.Println("It's a tie!")
		playerBoardToString(player1)
		playerBoardToString(computer)
	} else if player1.winner && !computer.winner {
		fmt.Printf("%s wins!\n", player1.name)
		playerBoardToString(player1)
	} else {
		fmt.Println("Computer wins!")
		playerBoardToString(computer)
	}
}
