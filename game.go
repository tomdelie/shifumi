package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	tm "github.com/buger/goterm"
)

// ROCK : rock
const ROCK string = "rock"

// PAPER : paper
const PAPER string = "paper"

// CISSORS : rock
const CISSORS string = "cissors"

var moves = []string{ROCK, PAPER, CISSORS}
var roundCounter int = 0
var p1Score int = 0
var p2Score int = 0
var p1Move int = 0
var p2Move int = 0

func main() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Flush()
	for {
		rand.Seed(time.Now().Unix())

		roundCounter++
		tm.Println(tm.Bold(logRoundStart(roundCounter)))
		tm.Flush()
		counter()

		p1Move = rand.Intn(3)
		p2Move = rand.Intn(3)
		fmt.Println(logPlayerMove(1, p1Move))
		fmt.Println(logPlayerMove(2, p2Move))

		time.Sleep(4 * time.Second)

		tm.Println(tm.Color(logTheWinner(p1Move, p2Move), tm.YELLOW))
		tm.Flush()

		time.Sleep(2 * time.Second)

		tm.Println(logPlayerScores(p1Score, p2Score))
		tm.Flush()

		if roundCounter >= 3 {
			return
		}

		time.Sleep(4 * time.Second)
	}
}

func whoIsTheWinner(p1Move string, p2Move string) int {
	if p1Move == p2Move {
		return 0
	}

	if (p1Move == ROCK && p2Move == CISSORS) ||
		(p1Move == PAPER && p2Move == ROCK) ||
		(p1Move == CISSORS && p2Move == PAPER) {
		return 1
	}

	return 2
}

func logPlayerScores(p1Score int, p2Score int) string {
	return "Scores: P1 (" + tm.Color(strconv.Itoa(p1Score), tm.BLUE) + ") | P2 (" + tm.Color(strconv.Itoa(p2Score), tm.BLUE) + ")"
}

func logTheWinner(p1Move int, p2Move int) string {
	if whoIsTheWinner(moves[p1Move], moves[p2Move]) == 1 {
		p1Score++
		return "Player 1 win."
	}

	if whoIsTheWinner(moves[p1Move], moves[p2Move]) == 2 {
		p2Score++
		return "Player 2 win."
	}

	return "It's a draw."
}

func logPlayerMove(playerID int, playerMove int) string {
	return "Player " + strconv.Itoa(playerID) + " play " + moves[playerMove] + "."
}

func logRoundStart(roundCounter int) string {
	return "========== Round " + strconv.Itoa(roundCounter) + " =========="
}

func counter() {
	for i := 3; i > 0; i-- {
		tm.Print(tm.Color(strconv.Itoa(i)+", ", tm.RED))
		time.Sleep(time.Second)
		tm.Flush()
	}
	time.Sleep(time.Second)
	tm.Println(tm.Color("Play!", tm.GREEN))
	tm.Flush()
}
