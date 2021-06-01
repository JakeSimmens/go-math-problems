package main

import (
	"testing"
)

func TestExecuteMenuSelection(t *testing.T) {
	//t.Log("GetMenuInput")
	problems := make([]problemSet, 1)
	problems[0] = problemSet{0, 0}
	score := scoreBoard{}
	score.init()

	var gotQuit bool

	gotQuit = executeMenuSelection("q", &problems, &score)
	if !gotQuit {
		t.Errorf("executeMenuSelection('q') = %v; Want TRUE to quit app", gotQuit)
	}

	gotQuit = executeMenuSelection("1", &problems, &score)
	if gotQuit {
		t.Errorf("executeMenuSelection('1') = %v; Want FALSE to continue app execution", gotQuit)
	}

	gotQuit = executeMenuSelection("2", &problems, &score)
	if gotQuit {
		t.Errorf("executeMenuSelection('2') = %v; Want FALSE to continue app execution", gotQuit)
	}

	gotQuit = executeMenuSelection("wrong", &problems, &score)
	if gotQuit {
		t.Errorf("executeMenuSelection('wrong') = %v; Want FALSE to continue app execution", gotQuit)
	}
}

func TestScoreBoard(t *testing.T) {
	testScore := scoreBoard{}
	testScore.init()

	if testScore.wins != 0 {
		t.Errorf("initializing testScore.wins = %v; Want initial value to be 0", testScore.wins)
	}
	if testScore.losses != 0 {
		t.Errorf("initializing testScore.losses = %v; Want initial value to be 0", testScore.losses)
	}
	if testScore.roundWins != 0 {
		t.Errorf("initializing testScore.roundWins = %v; Want initial value to be 0", testScore.roundWins)
	}
	if testScore.roundLosses != 0 {
		t.Errorf("initializing testScore.roundLosses = %v; Want initial value to be 0", testScore.roundLosses)
	}

	testScore.addWin()
	if testScore.wins != 1 {
		t.Errorf("initializing testScore.wins = %v; Want win value to be 1", testScore.wins)
	}
	if testScore.losses > 0 {
		t.Errorf("initializing testScore.losses = %v; Want loss value to be 0", testScore.losses)
	}
	if testScore.roundWins != 1 {
		t.Errorf("initializing testScore.roundWins = %v; Want win value to be 1", testScore.roundWins)
	}
	if testScore.roundLosses > 0 {
		t.Errorf("initializing testScore.roundLosses = %v; Want loss value to be 0", testScore.roundLosses)
	}

	testScore.addLoss()
	if testScore.wins > 1 {
		t.Errorf("initializing testScore.wins = %v; Want win value to be 1", testScore.wins)
	}
	if testScore.losses != 1 {
		t.Errorf("initializing testScore.losses = %v; Want loss value to be 1", testScore.losses)
	}
	if testScore.roundWins > 1 {
		t.Errorf("initializing testScore.roundWins = %v; Want win value to be 1", testScore.roundWins)
	}
	if testScore.roundLosses != 1 {
		t.Errorf("initializing testScore.roundLosses = %v; Want loss value to be 1", testScore.roundLosses)
	}

	testScore.resetRound()
	if testScore.wins < 1 {
		t.Errorf("initializing testScore.wins = %v; Want win value to be unchanged at 1", testScore.wins)
	}
	if testScore.losses < 1 {
		t.Errorf("initializing testScore.losses = %v; Want loss value to be unchanged at 1", testScore.losses)
	}
	if testScore.roundWins > 0 {
		t.Errorf("initializing testScore.roundWins = %v; Want win value reset to 0", testScore.roundWins)
	}
	if testScore.roundLosses > 0 {
		t.Errorf("initializing testScore.roundLosses = %v; Want loss value reset to 0", testScore.roundLosses)
	}
}
