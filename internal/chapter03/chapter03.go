package chapter03

import (
	"fmt"
	"time"
)

type WeekDay int

const (
	Pow0 = 1 << (iota * 2)
	Pow1
	Pow2
	Pow3
	Pow4
)

const (
	Monday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func Runner() {
	fmt.Println("*** Running Chapter 03 ***")

	ShowPowersOf4()
	ShowWeekdays()
	ConvertArrayToMap()
	ParseTime()
	ProcessTimeDate()
	ParseDate()
}

// ShowPowersOf4 shows constants that have the powers of number 4
func ShowPowersOf4() {
	fmt.Println("*** Running ShowPowersOf4 ***")
	fmt.Println(Pow0, Pow1, Pow2, Pow3, Pow4)
}

func (w WeekDay) String() string {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	if w < Monday || w > Sunday {
		return "This day is invalid"
	}
	return days[w]
}

// ShowWeekdays shows the weekdays correspondent to a number
func ShowWeekdays() {
	fmt.Println("*** Running ShowWeekdays ***")

	var monday WeekDay = Monday
	fmt.Println(monday.String())
}

// ConvertArrayToMap converts an existing array into a map
func ConvertArrayToMap() {
	fmt.Println("*** Running ConvertArrayToMap ***")

	a := [4]string{"This", "is", "a", "test"}
	m := make(map[int]string)

	for key, value := range a {
		m[key] = value
	}

	fmt.Println("result: ", m)
}

// ParseTime parses a time to the specified format
func ParseTime() {
	fmt.Println("*** Running ParseTime ***")

	t := time.Now()
	str := fmt.Sprint(t.Format("15:04:50"))
	result, err := time.Parse("15:04:50", str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("result: ", result)
}

// ProcessTimeDate process two date and time formats
func ProcessTimeDate() {
	fmt.Println("*** Running ProcessTimeDate ***")

	t := time.Now()
	fmt.Println("format 01: ", t.Format(time.ANSIC))
	fmt.Println("format 02: ", t.Format(time.Kitchen))
}

// ParseDate parses a date to the specified format
func ParseDate() {
	fmt.Println("*** Running ParseDate ***")

	t := time.Now()
	str := fmt.Sprint(t.Format("2006-01-02"))
	result, err := time.Parse("2006-01-02", str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("result: ", result)
}
