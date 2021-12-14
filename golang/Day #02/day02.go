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

func getData() string {
	dat, err := os.ReadFile("data.txt")
	check(err)
	return string(dat)
}

func getList(s string) []string {
	return strings.Split(s, "\n")
}

func getMultiplier(s string) Position {

	out := Position{0, 0}

	if s == "up" {
		out.y = -1
	} else if s == "down" {
		out.y = 1
	} else {
		out.x = 1
	}

	return out
}

func getAxisAndDirection(s string) Position {
	parts := strings.Split(s, " ")
	basePos := getMultiplier(parts[0])
	val, err := strconv.Atoi(parts[1])
	check(err)

	basePos.x = basePos.x * val
	basePos.y = basePos.y * val

	return basePos
}

type Position struct {
	x int
	y int
}

func doPart1(d []string) {
	pos := Position{x: 0, y: 0}

	for _, item := range d {
		mod := getAxisAndDirection(item)
		pos.x = pos.x + mod.x
		pos.y = pos.y + mod.y
	}

	fmt.Printf("Part1: Final value is %d!\n", (pos.x * pos.y))

}

func main() {
	data := getList(getData())
	doPart1(data)
	// doPart2(values)
}
