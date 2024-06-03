package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type game struct {
	calledNumbers map[string]int
}

func newGame() *game {
	game := game{calledNumbers: make(map[string]int)}
	return &game
}

type player struct {
	name   string
	wins   int
	board  [5][5]string
	winner bool
}

func newPlayer(name string) *player {
	player := player{name: name, wins: 0, board: [5][5]string{}, winner: false}
	return &player
}

func playerBoardToString(player *player) {
	fmt.Println("  B    I    N    G    O  ")
	fmt.Println("--------------------------------------")
	for i := 0; i < 5; i++ {
		fmt.Printf("| %s | %s | %s | %s | %s |\n",
			player.board[i][0], player.board[i][1], player.board[i][2], player.board[i][3], player.board[i][4])
	}
}

func playerWon(player *player) bool {
	for i := 0; i < 5; i++ {
		if player.board[i][0] == " X" && player.board[i][1] == " X" && player.board[i][2] == " X" &&
			player.board[i][3] == " X" && player.board[i][4] == " X" {
			for j := 0; j < 5; j++ {
				player.board[i][j] = " O"
			}
			return true
		}
		if player.board[0][i] == " X" && player.board[1][i] == " X" && player.board[2][i] == " X" &&
			player.board[3][i] == " X" && player.board[4][i] == " X" {
			for j := 0; j < 5; j++ {
				player.board[j][i] = " O"
			}
			return true
		}
	}
	if player.board[0][0] == " X" && player.board[1][1] == " X" && player.board[2][2] == " X" &&
		player.board[3][3] == " X" && player.board[4][4] == " X" {
		for j := 0; j < 5; j++ {
			player.board[j][j] = " O"
		}
		return true
	}
	return false
}

func (p *player) won() {
	p.winner = true
}

type Bingo int

const (
	B = iota
	I
	N
	G
	O
)

var bingoColumn = map[Bingo]string{
	B: "B",
	I: "I",
	N: "N",
	G: "G",
	O: "O",
}

func (bingoLetter Bingo) String() string {
	return bingoColumn[bingoLetter]
}

func newBoard() [5][5]string {
	intBoard := [5][5]string{}
	numbersInColumn := make(map[int]int)
	for i := 0; i < 5; i++ {
		j := 0
		for j < 5 {
			randNum := rand.Intn(100)
			_, exists := numbersInColumn[randNum]
			if !exists {
				numbersInColumn[randNum] = 1
				if randNum < 10 {
					intBoard[j][i] = "0" + strconv.Itoa(randNum)
				} else {
					intBoard[j][i] = strconv.Itoa(randNum)
				}
				j++
			}
		}
		clear(numbersInColumn)
	}
	intBoard[2][2] = " X"
	return intBoard
}

func caller(currentGame *game) (Bingo, int, int) {
	for {
		randNumber := rand.Intn(100)
		randBingoColumn := rand.Intn(5)
		randBingoLetter := Bingo(randBingoColumn)
		mapLetter := bingoColumn[randBingoLetter] + strconv.Itoa(randNumber)
		_, exists := currentGame.calledNumbers[mapLetter]
		if !exists {
			currentGame.calledNumbers[mapLetter] = 1
			return randBingoLetter, randBingoColumn, randNumber
		}
	}
}

func checkWon(player *player) bool {
	if playerWon(player) {
		player.wins++
	}
	return player.wins > 1
}
