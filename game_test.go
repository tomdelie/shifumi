package main

import (
	"testing"
)

// ! LOGGERS ================
func TestLogRoundStart(t *testing.T) {
	result := logRoundStart(1)
	var expectedResult string = "========== Round 1 =========="

	if result != expectedResult {
		t.Errorf("logRoundStart(1) FAILED, expected %v, but go %v", expectedResult, result)
	} else {
		t.Logf("logRoundStart(1) PASSED")
	}
}

func TestLogPlayerMove(t *testing.T) {
	result := logPlayerMove(1, 0)
	var expectedResult string = "Player 1 play rock."

	if result != expectedResult {
		t.Errorf("logPlayerMove(1, 0) FAILED, expected %v, but go %v", expectedResult, result)
	} else {
		t.Logf("logPlayerMove(1, 0) PASSED")
	}
}

func TestLogPlayerScores(t *testing.T) {
	result := logPlayerScores(3, 1)
	var expectedResult string = "Scores: P1 (3) | P2 (1)"

	if result != expectedResult {
		t.Errorf("logPlayerScores(3, 1) FAILED, expected %v, but go %v", expectedResult, result)
	} else {
		t.Logf("logPlayerScores(3, 1) PASSED")
	}
}

func TestLogTheWinner(t *testing.T) {
	// ! P2 win
	result1 := logTheWinner(0, 1)
	var expectedResult1 string = "Player 2 win."

	if result1 != expectedResult1 {
		t.Errorf("logTheWinner(0, 1) FAILED, expected %v, but go %v", expectedResult1, result1)
	} else {
		t.Logf("logTheWinner(0, 1) PASSED")
	}

	// ! P1 win
	result2 := logTheWinner(1, 0)
	var expectedResult2 string = "Player 1 win."

	if result2 != expectedResult2 {
		t.Errorf("logTheWinner(1, 0) FAILED, expected %v, but go %v", expectedResult2, result2)
	} else {
		t.Logf("logTheWinner(1, 0) PASSED")
	}

	// ! Draw
	result3 := logTheWinner(1, 1)
	var expectedResult3 string = "It's a draw."

	if result3 != expectedResult3 {
		t.Errorf("logTheWinner(1, 1) FAILED, expected %v, but go %v", expectedResult3, result3)
	} else {
		t.Logf("logTheWinner(1, 1) PASSED")
	}
}

func TestWhoIsTheWinner(t *testing.T) {
	// ! 0
	result1 := whoIsTheWinner(ROCK, ROCK)

	if result1 != 0 {
		t.Errorf("whoIsTheWinner(\"rock\", \"rock\") FAILED, expected %v, but go %v", 0, result1)
	} else {
		t.Logf("whoIsTheWinner(\"rock\", \"rock\") PASSED")
	}

	// ! 1
	result2 := whoIsTheWinner(PAPER, ROCK)

	if result2 != 1 {
		t.Errorf("whoIsTheWinner(\"paper\", \"rock\") FAILED, expected %v, but go %v", 1, result2)
	} else {
		t.Logf("whoIsTheWinner(\"paper\", \"rock\") PASSED")
	}

	// ! 2
	result3 := whoIsTheWinner(ROCK, PAPER)

	if result3 != 2 {
		t.Errorf("whoIsTheWinner(\"rock\", \"paper\") FAILED, expected %v, but go %v", 2, result3)
	} else {
		t.Logf("whoIsTheWinner(\"rock\", \"paper\") PASSED")
	}

}
