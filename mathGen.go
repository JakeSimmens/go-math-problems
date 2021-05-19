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

type colors struct {
	standard string
	green    string
	white    string
	red      string
}

func (c *colors) setColors() {
	if c.standard == "" {
		c.standard = "\033[0m"
	}
	if c.green == "" {
		c.green = "\033[32m"
	}
	if c.white == "" {
		c.white = "\033[37m"
	}
	if c.red == "" {
		c.red = "\033[31m"
	}
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
		cmd := getMenuInput()
		quit = executeMenuSelection(cmd, &problems)
	}
}

func showUserMenu() {
	color := colors{}
	color.setColors()
	//colorGreen := "\033[32m"
	colorReset := "\033[0m"

	fmt.Println(color.green, "MENU")
	fmt.Println("1. Addition +")
	fmt.Println("2. Subtraction -")
	fmt.Print("Q. Quit\n\n", colorReset)
}

func getMenuInput() string {
	var menuSelection string
	fmt.Print("Your Choice:  ")
	fmt.Scan(&menuSelection)
	return menuSelection
}

func executeMenuSelection(cmd string, numSets *[]problemSet) bool {
	if cmd == "q" || cmd == "Q" {
		return true
	}

	switch cmd {
	case "1":
		fmt.Println("Addition Time")
		runAddition(3, *numSets)
	case "2":
		fmt.Println("Subtraction Time")
		runSubtraction(3, *numSets)
	default:
		fmt.Print("\033[33m", "** Invalid option **\n\n", "\033[0m")
		//showUserMenu()
		//nextCmd := getMenuInput()
		//return executeMenuSelection(nextCmd, numSets)
	}
	return false
}

func runAddition(count int, numSets []problemSet) {
	changeToRed := "\033[31m"
	changeToWhite := "\033[37m"
	resetColor := "\033[0m"

	for i := 0; i < count; i++ {
		problemNum := rand.Intn(100)
		fmt.Printf(" %v + %v = ", numSets[problemNum].first, numSets[problemNum].second)
		correctAnswer := numSets[problemNum].first + numSets[problemNum].second
		var answer int
		fmt.Scan(&answer)
		if answer == correctAnswer {
			fmt.Print(changeToWhite, "CORRECT\n\n", resetColor)
		} else {
			fmt.Print(changeToRed, "Answer is: ", correctAnswer, "\n\n", resetColor)
		}
	}
}

func runSubtraction(count int, numSets []problemSet) {
	colorRed := "\033[31m"
	colorWhite := "\033[37m"
	colorReset := "\033[0m"

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
			fmt.Print(colorWhite, "CORRECT\n\n", colorReset)
		} else {
			fmt.Print(colorRed, "Answer is: ", correctAnswer, "\n\n", colorReset)
		}
	}
}
