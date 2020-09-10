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
	Move string
}

type playerStruct struct {
	score     int
	moveIndex int
}

var movesIndex = map[string]int{ROCK: 0, PAPER: 1, CISSORS: 2}
var moves = []string{ROCK, PAPER, CISSORS}

// ROUNDNUMBER : int
const ROUNDNUMBER int = 10

// ROCK : rock
const ROCK string = "rock"

// PAPER : paper
const PAPER string = "paper"

// CISSORS : rock
const CISSORS string = "cissors"

func handleGame(player *playerStruct, channel chan int) {
	for {
		player.moveIndex = <-channel
	}
}

func moveHandler(player *playerStruct, channel chan int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(moves[player.moveIndex]))
		case "POST":
			w.WriteHeader(http.StatusCreated)
			decoder := json.NewDecoder(r.Body)
			var p postStruct
			err := decoder.Decode(&p)
			if err != nil {
				panic(err)
			}
			if p.Move != "rock" && p.Move != "paper" && p.Move != "cissors" {
				w.Write([]byte(`{"message": "Please provide a valid move (rock, paper or cissors)."}`))
				break
			}

			tm.Println(tm.Color(tm.Bold("[API POST] Player will now play "+p.Move+"."), tm.RED))
			w.Write([]byte(`{"message": "You've successfully submit your move (` + p.Move + `)."}`))
			channel <- movesIndex[p.Move]
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "Can't find method requested"}`))
		}
	}
}

func api(player *playerStruct, channel chan int) {
	http.HandleFunc("/players/1/move", moveHandler(player, channel))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	var player = playerStruct{score: 0, moveIndex: 0}
	var bot = playerStruct{score: 0, moveIndex: 0}

	var roundCounter int = 0
	var channel chan int
	channel = make(chan int)
	go handleGame(&player, channel)
	go api(&player, channel)

	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Flush()
	for {
		rand.Seed(time.Now().Unix())

		roundCounter++
		tm.Println(tm.Bold(logRoundStart(roundCounter)))
		tm.Flush()
		counter()

		//p1Move = rand.Intn(3)
		bot.moveIndex = rand.Intn(3)
		fmt.Println(logPlayerMove("Player", &player))
		fmt.Println(logPlayerMove("Bot", &bot))

		time.Sleep(4 * time.Second)

		tm.Println(tm.Color(logTheWinner(&player, &bot), tm.YELLOW))
		tm.Flush()

		time.Sleep(2 * time.Second)

		tm.Println(logPlayerScores(&player, &bot))
		tm.Flush()

		if roundCounter >= ROUNDNUMBER {
			return
		}

		time.Sleep(4 * time.Second)
	}
}

func whoIsTheWinner(player *playerStruct, bot *playerStruct) int {
	if player.moveIndex == bot.moveIndex {
		return 0
	}

	if (moves[player.moveIndex] == ROCK && moves[bot.moveIndex] == CISSORS) ||
		(moves[player.moveIndex] == PAPER && moves[bot.moveIndex] == ROCK) ||
		(moves[player.moveIndex] == CISSORS && moves[bot.moveIndex] == PAPER) {
		return 1
	}

	return 2
}

func logPlayerScores(player *playerStruct, bot *playerStruct) string {
	return "Scores: Player (" + strconv.Itoa(player.score) + ") | Bot (" + strconv.Itoa(bot.score) + ")"
}

func logTheWinner(player *playerStruct, bot *playerStruct) string {
	if whoIsTheWinner(player, bot) == 1 {
		player.score++
		return "Player win."
	}

	if whoIsTheWinner(player, bot) == 2 {
		bot.score++
		return "Bot win."
	}

	return "It's a draw."
}

func logPlayerMove(playerType string, player *playerStruct) string {
	return playerType + " play " + moves[player.moveIndex] + "."
}

func logRoundStart(roundCounter int) string {
	return "========== Round " + strconv.Itoa(roundCounter) + " =========="
}

func counter() {
	for i := 5; i > 0; i-- {
		tm.Print(tm.Color(strconv.Itoa(i)+", ", tm.RED))
		time.Sleep(time.Second)
		tm.Flush()
	}
	time.Sleep(time.Second)
	tm.Println(tm.Color("Play!", tm.GREEN))
	tm.Flush()
}
