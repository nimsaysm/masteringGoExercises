package main

import (
	"fmt"
	"github.com/nimsaysm/masteringGoExercises/internal/chapter01"
	"github.com/nimsaysm/masteringGoExercises/internal/chapter02"
	"github.com/nimsaysm/masteringGoExercises/internal/chapter03"
	"github.com/nimsaysm/masteringGoExercises/internal/chapter04"
)

func RunExercises() {
	fmt.Println("Choose which chapter run...")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		chapter01.Runner()
	case 2:
		chapter02.Runner()
	case 3:
		chapter03.Runner()
	case 4:
		chapter04.Runner()
	default:
		fmt.Println("The chapter does not exist.")
	}
}
