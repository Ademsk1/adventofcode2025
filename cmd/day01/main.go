package day01

import (
	"log"
	"strconv"
	"strings"

	printer "github.com/Ademsk1/adventofcode2025/util/print"
	"github.com/Ademsk1/adventofcode2025/util/read"
)

func Day01(part int, useExample bool) {

	content := read.Read("day01", useExample)
	data := read.LinesAsArray(content, "\n")
	switch part {
	case 1:
		part1(data)
	case 2:
		part2(data)
	default:
		part1(data)
		part2(data)
	}

}

func ParseTurn(line string, rotationSize *int, sign *int) {
	if line[0] == 'L' {
		*sign = -1
	} else {
		*sign = 1
	}
	r, err := strconv.ParseInt(strings.TrimSpace(line[1:]), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	*rotationSize = int(r)

}

func part1(data []string) {
	zeroCounter := 0
	dial := 50
	sign := -1
	rotationSize := 0

	for _, line := range data {
		if line == "" {
			continue
		}
		ParseTurn(line, &rotationSize, &sign)

		dial += (int(rotationSize) * sign)

		if dial%100 == 0 {
			zeroCounter++
		}

	}
	printer.PrettyPrintResult("1", "1", zeroCounter)

}

func part2(data []string) {
	var rotationSize int
	var sign int
	dial := 50
	zeroCounter := 0
	for _, line := range data {
		ParseTurn(line, &rotationSize, &sign)
		for i := 0; i < rotationSize; i++ {
			dial += sign
			if dial%100 == 0 {
				dial = 0
				zeroCounter++
			}
			if dial < 0 {
				dial = 100 + dial
			}

		}
	}
	printer.PrettyPrintResult("2", "2", zeroCounter)

}
