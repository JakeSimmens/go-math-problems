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
	fmt.Println("Welcome to the Math Generator\n")

	//generate random addition problems between 2 numbers
	problems := make([]problemSet, 100)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			problems[i*10+j] = problemSet{i, j}
		}
	}

	//problemUsed := make([]bool, 100)
	rand.Seed(time.Now().UnixNano())

	// for i := 0; i < 10; i++ {
	// 	problemNum := rand.Intn(100)
	// 	fmt.Println(problems[problemNum])
	// }

	for i := 0; i < 5; i++ {
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

	//fmt.Println(problems)
	//display problem
	//get user input
	//check answer
	//display correct or wrong
	//show next problem

	//track problems that are used and don't repeat
}
