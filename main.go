package main

import (
	"fmt"
	"os"

	"github.com/rahul-choudhury/aoc-2025/day01"
)

func main() {
	var args = os.Args
	var day string

	if len(args) == 2 {
		day = args[1]
	}

	switch day {
	case "01":
		fmt.Println(day01.Part1())
	default:
		fmt.Println("No Args passed.")
	}
}
