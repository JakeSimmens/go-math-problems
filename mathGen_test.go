package main

import (
	"testing"
)

func TestExecuteMenuSelection(t *testing.T) {
	//t.Log("GetMenuInput")
	problems := make([]problemSet, 1)
	problems[0] = problemSet{0, 0}

	got := executeMenuSelection("q", &problems)
	if !got {
		t.Errorf("getMenuInput() = %v; want true to quit app", got)
	}
}
