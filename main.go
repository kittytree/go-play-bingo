package main

import (
	"fmt"
	"time"
)

func main() {

	playerName := getPlayerName()

	computer := newPlayer("computer")
	player1 := newPlayer(playerName)

	computer.board = newBoard()
	game := newGame()
	playerBoardChooser(player1)

	fmt.Printf("%s board:\n", player1.name)
	playerBoardToString(player1)
	fmt.Println("Computer board:")
	playerBoardToString(computer)

	for !player1.winner && !computer.winner {
		letter, column, number := caller(game)
		fmt.Printf("Does anyone have: %s,%v?\n", letter, number)
		time.Sleep(500 * time.Millisecond)
		checkIfCallerMatch(player1, column, number)
		checkIfCallerMatch(computer, column, number)
	}

	announceWinner(player1, computer)
}
