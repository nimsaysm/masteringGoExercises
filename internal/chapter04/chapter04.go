package chapter04

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

type ipEntry struct {
	ip       string
	countIPs int
}

func Runner() {
	fmt.Println("*** Running Chapter 04 ***")

	CalculateSquareRoots()
	FindBetween200And400()
}

// CalculateSquareRoots calculates square roots with precision
func CalculateSquareRoots() {
	fmt.Println("*** Running CalculateSquareRoots ***")
	var result *big.Float

	if len(os.Args) > 1 {
		value := os.Args[1]
		result = calculate(value)
	} else {
		fmt.Println("Does not have values to calculate square root")
		return
	}

	fmt.Println("result: ", result)
}

func calculate(value string) *big.Float {
	num := new(big.Float)
	_, status := num.SetString(value)
	if status != true {
		fmt.Println("Invalid number")
		return nil
	}

	precision := uint(100)
	tryValue := new(big.Float).SetPrec(precision).SetFloat64(1.0)
	threshold := new(big.Float).SetPrec(precision).SetFloat64(1e-50) // Precision threshold
	delta := new(big.Float)
	two := big.NewFloat(2)

	for {
		// Newton-Raphson iteration: tryValue = (tryValue + num/tryValue) / 2
		temp := new(big.Float).Quo(num, tryValue)
		temp.Add(temp, tryValue)
		temp.Quo(temp, two)

		delta.Sub(temp, tryValue).Abs(delta)
		if delta.Cmp(threshold) < 0 {
			return temp
		}

		tryValue.Set(temp)
	}
}

// FindBetween200And400 finds all integers between 200 and 400 and prints it
func FindBetween200And400() {
	var matchesList []int
	pattern := `\b(2[0-9]{2}|3[0-9]{2}|400)\b`
	regex := regexp.MustCompile(pattern)
	fmt.Println("Enter numbers or STOP:")
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		line := reader.Text()
		if line == "STOP" {
			break
		}

		matches := regex.FindAllString(line, -1)
		for _, m := range matches {
			num, err := strconv.Atoi(m)
			if err != nil {
				fmt.Println("got error: ", err.Error())
				return
			}

			if num >= 200 && num <= 400 {
				matchesList = append(matchesList, num)
			}
		}
	}

	fmt.Println("matches: ", matchesList)
}
