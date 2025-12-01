package day01

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/rahul-choudhury/aoc-2025/utils"
)

func Part1() int {
	data := utils.ReadFile("input.txt")
	dataToArr := strings.Split(data, "\n")

	regexCond := regexp.MustCompile(`([LR])(\d+)`)

	dialPosition := 50
	zeroCount := 0

	for _, v := range dataToArr {
		parts := regexCond.FindStringSubmatch(v)
		if parts == nil {
			continue
		}
		num, _ := strconv.Atoi(parts[2])

		if parts[1] == "L" {
			dialPosition = ((dialPosition-num)%100 + 100) % 100
		} else {
			dialPosition = (dialPosition + num) % 100
		}

		if dialPosition == 0 {
			zeroCount++
		}

	}

	return zeroCount
}

func Part2() int {
	data := utils.ReadFile("input.txt")
	dataToArr := strings.Split(data, "\n")

	regexCond := regexp.MustCompile(`([LR])(\d+)`)

	dialPosition := 50
	zeroCount := 0

	for _, v := range dataToArr {
		parts := regexCond.FindStringSubmatch(v)
		if parts == nil {
			continue
		}
		num, _ := strconv.Atoi(parts[2])

		if parts[1] == "L" {
			newPosition := ((dialPosition-num)%100 + 100) % 100
			// NOTE: int division rounds to 0. so values like 0.6 or -1.2 will round up to 0.
			// hence the large bias to convert the negative numbers to large positive numbers.
			// idk how good this method is. it's either this or run a loop from 1 till num
			// and simulate crossing of 0. i should get back to this later.
			bias := 1000000
			zeroCount += (dialPosition-1+bias)/100 - (dialPosition-num-1+bias)/100
			dialPosition = newPosition
		} else {
			newPosition := (dialPosition + num) % 100
			zeroCount += (dialPosition + num) / 100
			dialPosition = newPosition
		}
	}

	return zeroCount
}
