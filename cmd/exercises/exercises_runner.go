package main

import (
	"fmt"
	"github.com/nimsaysm/masteringGoExercises/internal/chapter01"
)

func RunExercises() {
	fmt.Println("Choose which chapter run...")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		chapter01.Runner()
	default:
		fmt.Println("The chapter does not exist.")
	}
}
