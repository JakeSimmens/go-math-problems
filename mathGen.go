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
	fmt.Print("Welcome to the Math Generator\n\n")

	showUserMenu()
	getUserMenuSelection()

	//problemUsed := make([]bool, 100)
	rand.Seed(time.Now().UnixNano())
	additionProblems(3)

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
	fmt.Println("Enter a number")
	fmt.Println("1. Addition +")
	fmt.Println("2. Subtraction -")
	fmt.Print("Q. to Quit\n\n")
	fmt.Print("Your Choice:  ")

}

func getUserMenuSelection() {
	var menuSelection string
	fmt.Scan(&menuSelection)

	switch menuSelection {
	case "Q":
		return
	case "1":
		fmt.Println("Addition Time")
	case "2":
		fmt.Println("Subtraction Time")
	default:
		fmt.Print("Please make a selection:  ")
		fmt.Scan(&menuSelection)
	}
}

func additionProblems(count int) {
	//generate random addition problems between 2 numbers
	problems := make([]problemSet, 100)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			problems[i*10+j] = problemSet{i, j}
		}
	}

	for i := 0; i < count; i++ {
		problemNum := rand.Intn(100)
		fmt.Printf(" %v + %v = ", problems[problemNum].first, problems[problemNum].second)
		correctAnswer := problems[problemNum].first + problems[problemNum].second
		var answer int
		fmt.Scan(&answer)
		if answer == correctAnswer {
			fmt.Print("CORRECT\n\n")
		} else {
			fmt.Print("Answer is: ", correctAnswer, "\n\n")
		}
	}
}
