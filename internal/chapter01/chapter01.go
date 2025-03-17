package chapter01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Runner() {
	fmt.Println("*** Running Chapter 01 ***")

	SumNumbers()
	AverageNumbers()
	ReadInts()
}

// SumNumbers finds the sum of all of its numeric command-line arguments
func SumNumbers() {
	fmt.Println("*** Running SumNumbers ***")
	cliContent := os.Args
	if len(cliContent) == 1 {
		fmt.Println("There is any number to sum. Ending the program...")
		return
	}

	var numbers []int
	var sum int
	for _, value := range cliContent[1:] {
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("err: ", err.Error())
			continue
		}
		numbers = append(numbers, number)
		sum += number
	}

	fmt.Println("Numbers: ", numbers)
	fmt.Println("Result: ", sum)
}

// AverageNumbers finds the average value of all of its float command-line arguments
func AverageNumbers() {
	fmt.Println("*** Running AverageNumbers ***")
	cliContent := os.Args
	if len(cliContent) == 1 {
		fmt.Println("There is any number to sum. Ending the program...")
		return
	}

	var numbers []float64
	var total, average float64
	for _, value := range cliContent[1:] {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("err: ", err.Error())
			continue
		}
		numbers = append(numbers, number)
		total += number
	}

	elementsQuantity := len(numbers)
	average = total / float64(elementsQuantity)

	fmt.Println("Numbers: ", numbers)
	fmt.Printf("Result: %.2f\n", average)
}

// ReadInts keeps reading integers until it gets the word STOP as input
func ReadInts() {
	fmt.Println("*** Running ReadInts ***")

	for {
		fmt.Println("Input an int or close the program typing STOP")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		if input == "STOP" {
			return
		}
		if _, err := strconv.Atoi(input); err != nil {
			fmt.Println("Unexpected input. Ending program...")
			return
		}

	}
}
