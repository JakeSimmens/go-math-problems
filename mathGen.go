package main

import "fmt"

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

	fmt.Println(problems)
	//display problem
	//get user input
	//check answer
	//display correct or wrong
	//show next problem

	//track problems that are used and don't repeat
}
