package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var p1Move = "rock"

type postStruct struct {
	move string
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(p1Move))
	case "POST":
		//w.WriteHeader(http.StatusCreated)
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
		p1Move = s.move
		w.Write([]byte(`{"message": "You've successfully submit your move (` + s.move + `)."}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func main() {
	http.HandleFunc("/players/1/move", apiResponse)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
