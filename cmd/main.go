package main

import (
	"github.com/Ademsk1/adventofcode2025/cmd/day01"
	"github.com/Ademsk1/adventofcode2025/cmd/day02"
	"github.com/Ademsk1/adventofcode2025/cmd/day03"
	"github.com/Ademsk1/adventofcode2025/util/flags"
)

func main() {
	day, part, useExample := flags.ParseFlags()
	switch day {
	case 1:
		day01.Day01(part, useExample)
	case 2:
		day02.Day02(part, useExample)
	case 3:
		day03.Day03(part, useExample)
	}
}
