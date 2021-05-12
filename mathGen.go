package main

import (
	"fmt"
	"math/rand"
	"time"
)

type problemSet struct {
	first  int
	second int
}

func main() {

	problems := make([]problemSet, 100)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			problems[i*10+j] = problemSet{i, j}
		}
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Print("Welcome to the Math Generator\n\n")

	quit := false
	for !quit {
		showUserMenu()
		quit = executeUserSelection(&problems)
	}

	//problemUsed := make([]bool, 100)

	//additionProblems(3)

	//fmt.Println(problems)
	//display problem
	//get user input
	//check answer
	//display correct or wrong
	//show next problem

	//track problems that are used and don't repeat
}

func showUserMenu() {

	fmt.Println("MENU")
	fmt.Println("1. Addition +")
	fmt.Println("2. Subtraction -")
	fmt.Print("Q. Quit\n\n")
}

func executeUserSelection(numSets *[]problemSet) bool {
	var menuSelection string
	fmt.Print("Your Choice:  ")
	fmt.Scan(&menuSelection)
	if menuSelection == "q" || menuSelection == "Q" {
		return true
	}

	switch menuSelection {
	case "1":
		fmt.Println("Addition Time")
		additionProblems(3, *numSets)
	case "2":
		fmt.Println("Subtraction Time")
	default:
		fmt.Print("Invalid option.\n\n")
		showUserMenu()
		return executeUserSelection(numSets)
	}
	return false
}

func additionProblems(count int, numSets []problemSet) {
	//generate random addition problems between 2 numbers
	// problems := make([]problemSet, 100)
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		problems[i*10+j] = problemSet{i, j}
	// 	}
	// }

	for i := 0; i < count; i++ {
		problemNum := rand.Intn(100)
		fmt.Printf(" %v + %v = ", numSets[problemNum].first, numSets[problemNum].second)
		correctAnswer := numSets[problemNum].first + numSets[problemNum].second
		var answer int
		fmt.Scan(&answer)
		if answer == correctAnswer {
			fmt.Print("CORRECT\n\n")
		} else {
			fmt.Print("Answer is: ", correctAnswer, "\n\n")
		}
	}
}
