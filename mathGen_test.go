package main

import (
	"testing"
)

func TestExecuteMenuSelection(t *testing.T) {
	//t.Log("GetMenuInput")
	problems := make([]problemSet, 1)
	problems[0] = problemSet{0, 0}

	gotQuit := executeMenuSelection("q", &problems)
	if !gotQuit {
		t.Errorf("executeMenueSelection() = %v; want true to quit app", gotQuit)
	}

	gotQuit = executeMenuSelection("1", &problems)
	if gotQuit {
		t.Errorf("executeMenueSelection() = %v; want false to continue app running", gotQuit)
	}
}
