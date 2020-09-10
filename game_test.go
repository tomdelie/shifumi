package main

import (
	"testing"
)

// TestLogRoundStart : test
func TestLogRoundStart(t *testing.T) {
	result := logRoundStart(1)
	var expectedResult string = "========== Round 1 =========="

	if result != expectedResult {
		t.Errorf("logRoundStart(1) FAILED, expected %v, but go %v", expectedResult, result)
	} else {
		t.Logf("logRoundStart(1) PASSED")
	}
}

// TestLogPlayerMove : test
func TestLogPlayerMove(t *testing.T) {
	var player = playerStruct{0, 0}
	result := logPlayerMove("Player", &player)
	var expectedResult string = "Player play rock."

	if result != expectedResult {
		t.Errorf("logPlayerMove(\"Player\", &player) FAILED, expected %v, but go %v", expectedResult, result)
	} else {
		t.Logf("logPlayerMove(\"Player\", &player) PASSED")
	}
}

// TestLogPlayerScores : test
func TestLogPlayerScores(t *testing.T) {
	var player = playerStruct{3, 0}
	var bot = playerStruct{1, 0}

	result := logPlayerScores(&player, &bot)
	var expectedResult string = "Scores: Player (3) | Bot (1)"

	if result != expectedResult {
		t.Errorf("logPlayerScores(&player, &bot) FAILED, expected %v, but go %v", expectedResult, result)
	} else {
		t.Logf("logPlayerScores(&player, &bot) PASSED")
	}
}

// TestLogTheWinner : test
func TestLogTheWinner(t *testing.T) {
	var player = playerStruct{0, 0}
	var bot = playerStruct{0, 1}

	// ! P2 win
	result1 := logTheWinner(&player, &bot)
	var expectedResult1 string = "Bot win."

	if result1 != expectedResult1 {
		t.Errorf("logTheWinner(&player, &bot) FAILED, expected %v, but go %v", expectedResult1, result1)
	} else {
		t.Logf("logTheWinner(&player, &bot) PASSED")
	}

	// ! P1 win
	player.moveIndex = 2
	result2 := logTheWinner(&player, &bot)
	var expectedResult2 string = "Player win."

	if result2 != expectedResult2 {
		t.Errorf("logTheWinner(&player, &bot) FAILED, expected %v, but go %v", expectedResult2, result2)
	} else {
		t.Logf("logTheWinner(&player, &bot) PASSED")
	}

	// ! Draw
	player.moveIndex = 1
	result3 := logTheWinner(&player, &bot)
	var expectedResult3 string = "It's a draw."

	if result3 != expectedResult3 {
		t.Errorf("logTheWinner(&player, &bot) FAILED, expected %v, but go %v", expectedResult3, result3)
	} else {
		t.Logf("logTheWinner(&player, &bot) PASSED")
	}
}

// TestWhoIsTheWinner : test
func TestWhoIsTheWinner(t *testing.T) {
	var player = playerStruct{0, 0}
	var bot = playerStruct{0, 0}

	// ! 0
	result1 := whoIsTheWinner(&player, &bot)

	if result1 != 0 {
		t.Errorf("whoIsTheWinner(\"rock\", \"rock\") FAILED, expected %v, but go %v", 0, result1)
	} else {
		t.Logf("whoIsTheWinner(\"rock\", \"rock\") PASSED")
	}

	// ! 1
	player.moveIndex = 1
	result2 := whoIsTheWinner(&player, &bot)

	if result2 != 1 {
		t.Errorf("whoIsTheWinner(\"paper\", \"rock\") FAILED, expected %v, but go %v", 1, result2)
	} else {
		t.Logf("whoIsTheWinner(\"paper\", \"rock\") PASSED")
	}

	// ! 2
	player.moveIndex = 0
	bot.moveIndex = 1
	result3 := whoIsTheWinner(&player, &bot)

	if result3 != 2 {
		t.Errorf("whoIsTheWinner(\"rock\", \"paper\") FAILED, expected %v, but go %v", 2, result3)
	} else {
		t.Logf("whoIsTheWinner(\"rock\", \"paper\") PASSED")
	}
}
