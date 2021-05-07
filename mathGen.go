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
	fmt.Println("Welcome to the Math Generator")

	//generate random addition problems between 2 numbers
	problems := make([]problemSet, 100)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			problems[i*10+j] = problemSet{i, j}
		}
	}

	//problemUsed := make([]bool, 100)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		problemNum := rand.Intn(100)
		fmt.Println(problems[problemNum])
	}

	//fmt.Println(problems)
	//display problem
	//get user input
	//check answer
	//display correct or wrong
	//show next problem

	//track problems that are used and don't repeat
}
