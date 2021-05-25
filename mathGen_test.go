package main

import (
	"testing"
)

func TestExecuteMenuSelection(t *testing.T) {
	//t.Log("GetMenuInput")
	problems := make([]problemSet, 1)
	problems[0] = problemSet{0, 0}

	var gotQuit bool

	gotQuit = executeMenuSelection("q", &problems)
	if !gotQuit {
		t.Errorf("executeMenuSelection('q') = %v; Want TRUE to quit app", gotQuit)
	}

	gotQuit = executeMenuSelection("1", &problems)
	if gotQuit {
		t.Errorf("executeMenuSelection('1') = %v; Want FALSE to continue app execution", gotQuit)
	}

	gotQuit = executeMenuSelection("2", &problems)
	if gotQuit {
		t.Errorf("executeMenuSelection('2') = %v; Want FALSE to continue app execution", gotQuit)
	}

	gotQuit = executeMenuSelection("wrong", &problems)
	if gotQuit {
		t.Errorf("executeMenuSelection('wrong') = %v; Want FALSE to continue app execution", gotQuit)
	}
}
