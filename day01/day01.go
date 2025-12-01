package day01

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/rahul-choudhury/aoc-2025/utils"
)

type instruction struct {
	direction string
	num       int
}

func parseInput() []instruction {
	data := utils.ReadFile("input.txt")
	dataToArr := strings.Split(data, "\n")
	regexCond := regexp.MustCompile(`([LR])(\d+)`)

	var instructions []instruction

	for _, v := range dataToArr {
		parts := regexCond.FindStringSubmatch(v)
		if parts == nil {
			continue
		}
		num, _ := strconv.Atoi(parts[2])
		instructions = append(instructions, instruction{
			direction: parts[1],
			num:       num,
		})
	}

	return instructions
}

func moveDial(pos, num int, direction string) int {
	if direction == "L" {
		return ((pos-num)%100 + 100) % 100
	} else {
		return (pos + num) % 100
	}
}

func Part1() int {
	instructions := parseInput()
	dialPosition := 50
	zeroCount := 0

	for _, inst := range instructions {
		dialPosition = moveDial(dialPosition, inst.num, inst.direction)
		if dialPosition == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func Part2() int {
	instructions := parseInput()
	dialPosition := 50
	zeroCount := 0

	for _, inst := range instructions {
		if inst.direction == "L" {
			newPosition := moveDial(dialPosition, inst.num, inst.direction)
			// NOTE: int division rounds to 0. so values like 0.6 or -1.2 will round up to 0.
			// hence the large bias to convert the negative numbers to large positive numbers.
			// idk how good this method is. it's either this or run a loop from 1 till num
			// and simulate crossing of 0. i should get back to this later.
			bias := 1000000
			zeroCount += (dialPosition-1+bias)/100 - (dialPosition-inst.num-1+bias)/100
			dialPosition = newPosition
		} else {
			newPosition := moveDial(dialPosition, inst.num, inst.direction)
			zeroCount += (dialPosition + inst.num) / 100
			dialPosition = newPosition
		}
	}

	return zeroCount
}
