package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getList(s string) []string {
	return strings.Split(s, "\n")
}

func convertStringToInt(s string) int {
	val, err := strconv.Atoi(s)
	check(err)
	return val
}

func getData() string {
	dat, err := os.ReadFile("sonar-readings.txt")
	check(err)
	return string(dat)
}

func convertDataToIntegers(s []string) []int {
	out := []int{}
	for i := 0; i < len(s); i++ {
		out = append(
			out,
			convertStringToInt(s[i]),
		)
	}
	return out
}

func isHigher(prev int, cur int) int {
	if cur > prev {
		return 1
	}
	return 0
}

func addElements(e []int) int {
	sum := 0
	for _, item := range e {
		sum = sum + item
	}
	return sum
}

func doPart1(values []int) {
	numOfReadings := 0

	for i := 1; i < len(values); i++ {
		cur := values[i]
		prev := values[i-1]
		numOfReadings = numOfReadings + isHigher(prev, cur)

	}

	fmt.Printf("Part 1: Found %v increased readings.\n", numOfReadings)
}

func doPart2(values []int) {

	// A sonar reading can never be negative
	// and I'm not sure if nullable types are a thing
	// in golang :D
	prev := -1
	cur := -1
	increases := 0

	for i := 3; i <= len(values); i++ {
		prev = cur

		sliceOfValues := values[(i - 3):i]
		cur = addElements(sliceOfValues)

		if prev >= 0 && cur > prev {
			increases = increases + 1
		}
	}

	fmt.Printf("Part 2: Found %v increased averages.\n", increases)
}

func main() {
	data := getList(getData())
	values := convertDataToIntegers(data)

	doPart1(values)
	doPart2(values)
}
