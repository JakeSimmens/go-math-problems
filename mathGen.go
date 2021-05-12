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
		subtractionProblems(3, *numSets)
	default:
		fmt.Print("Invalid option.\n\n")
		showUserMenu()
		return executeUserSelection(numSets)
	}
	return false
}

func additionProblems(count int, numSets []problemSet) {

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

func subtractionProblems(count int, numSets []problemSet) {

	for i := 0; i < count; i++ {
		var firstNum int
		var secondNum int

		problemNum := rand.Intn(100)
		if numSets[problemNum].second > numSets[problemNum].first {
			firstNum = numSets[problemNum].second
			secondNum = numSets[problemNum].first
		} else {
			firstNum = numSets[problemNum].first
			secondNum = numSets[problemNum].second
		}

		fmt.Printf(" %v - %v = ", firstNum, secondNum)
		correctAnswer := firstNum - secondNum
		var answer int
		fmt.Scan(&answer)
		if answer == correctAnswer {
			fmt.Print("CORRECT\n\n")
		} else {
			fmt.Print("Answer is: ", correctAnswer, "\n\n")
		}
	}
}
