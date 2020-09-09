package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	tm "github.com/buger/goterm"
)

type postStruct struct {
	move string
}

// ROCK : rock
const ROCK string = "rock"

// PAPER : paper
const PAPER string = "paper"

// CISSORS : rock
const CISSORS string = "cissors"

var movesIndex = map[string]int{ROCK: 0, PAPER: 1, CISSORS: 2}
var moves = []string{ROCK, PAPER, CISSORS}
var roundCounter int = 0
var p1Score int = 0
var p2Score int = 0
var p1Move int = 0
var p2Move int = 0

func api() {
	http.HandleFunc("/players/1/move", apiResponse)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	go api()
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

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(moves[p1Move]))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		decoder := json.NewDecoder(r.Body)
		var s postStruct
		err := decoder.Decode(&s)
		if err != nil {
			panic(err)
		}
		fmt.Print(s)
		if s.move != "rock" && s.move != "paper" && s.move != "cissors" {
			w.Write([]byte(`{"message": "Please provide a valid move (rock, paper or cissors)."}`))
			break
		}
		p1Move = movesIndex[s.move] // ! should call a goroutine
		w.Write([]byte(`{"message": "You've successfully submit your move (` + s.move + `)."}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
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
