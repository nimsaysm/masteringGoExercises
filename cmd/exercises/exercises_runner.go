package main

import "fmt"

func RunExercises() {
	fmt.Println("Choose which chapter run...")
	var input int
	fmt.Scan(&input)
	switch input {
	default:
		fmt.Println("The chapter does not exist.")
	}
}
