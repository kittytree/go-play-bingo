package main

import (
	"math/rand"
	"strconv"
)

type game struct {
	players []*player
}

func newGame() *game {
	return &game{}
}

func addPlayerToGame(g *game, p *player) {
	g.players = append(g.players, p)
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
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			intBoard[i][j] = strconv.Itoa(rand.Intn(100))
		}
	}
	intBoard[2][2] = "x"
	return intBoard
}

func caller() (Bingo, int, int) {
	randNumber := rand.Intn(100)
	randBingoColumn := rand.Intn(4)
	randBingoLetter := Bingo(randBingoColumn)
	return randBingoLetter, randBingoColumn, randNumber

}
