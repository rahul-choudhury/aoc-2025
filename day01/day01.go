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

func Part2() {
}
