package main

import (
	"fmt"
	"math"
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
	yellow   string
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
	if c.yellow == "" {
		c.yellow = "\033[33m"
	}
}

type scoreBoard struct {
	wins        int
	losses      int
	roundWins   int
	roundLosses int
}

func (score *scoreBoard) init() {
	score.wins = 0
	score.losses = 0
	score.roundWins = 0
	score.roundLosses = 0
}

func (score *scoreBoard) addWin() {
	score.wins++
	fmt.Println("wins: ", score.wins)
	fmt.Println("losses: ", score.losses)
}

func (score *scoreBoard) addLoss() {
	score.losses++
	fmt.Println("wins: ", score.wins)
	fmt.Println("losses: ", score.losses)
}

func (score *scoreBoard) showScore() {
	totalProblems := score.wins + score.losses
	fmt.Println("Total Wins:  ", score.wins)
	fmt.Println("Total Losses:  ", score.losses)
	fmt.Println("Winning %:  ", math.Floor((float64(score.wins)/float64(totalProblems))*100))
}

func main() {

	score := scoreBoard{}
	score.init()

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
		quit = executeMenuSelection(cmd, &problems, &score)
		score.showScore()

	}
}

func showUserMenu() {
	color := colors{}
	color.setColors()

	fmt.Println(color.green, "  MENU\n", color.standard)
	fmt.Println("1. Addition +")
	fmt.Println("2. Subtraction -")
	fmt.Print("Q. Quit\n\n")
}

func getMenuInput() string {
	var menuSelection string
	fmt.Print("Your Choice:  ")
	fmt.Scan(&menuSelection)
	return menuSelection
}

func executeMenuSelection(cmd string, numSets *[]problemSet, score *scoreBoard) bool {
	if cmd == "q" || cmd == "Q" {
		return true
	}
	color := colors{}
	color.setColors()

	switch cmd {
	case "1":
		fmt.Println("Addition Time")
		runAddition(3, *numSets, score)
	case "2":
		fmt.Println("Subtraction Time")
		runSubtraction(3, *numSets, score)
	default:
		fmt.Print(color.yellow, "** Invalid option **\n\n", color.standard)
	}
	return false
}

func runAddition(count int, numSets []problemSet, score *scoreBoard) {
	color := colors{}
	color.setColors()

	for i := 0; i < count; i++ {
		problemNum := rand.Intn(len(numSets))
		fmt.Printf(" %v + %v = ", numSets[problemNum].first, numSets[problemNum].second)
		correctAnswer := numSets[problemNum].first + numSets[problemNum].second
		var answer int
		fmt.Scan(&answer)
		if answer == correctAnswer {
			score.addWin()
			fmt.Print(color.white, "CORRECT\n\n", color.standard)
		} else {
			score.addLoss()
			fmt.Print(color.red, "Answer is: ", correctAnswer, "\n\n", color.standard)
		}
	}
}

func runSubtraction(count int, numSets []problemSet, score *scoreBoard) {
	color := colors{}
	color.setColors()

	for i := 0; i < count; i++ {
		var firstNum int
		var secondNum int

		problemNum := rand.Intn(len(numSets))
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
			score.addWin()
			fmt.Print(color.white, "CORRECT\n\n", color.standard)
		} else {
			score.addLoss()
			fmt.Print(color.red, "Answer is: ", correctAnswer, "\n\n", color.standard)
		}
	}
}
