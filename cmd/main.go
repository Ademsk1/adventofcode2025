package main

import (
	"fmt"

	"github.com/Ademsk1/adventofcode2025/cmd/day01"
	"github.com/Ademsk1/adventofcode2025/cmd/day02"
	"github.com/Ademsk1/adventofcode2025/util/flags"
)

func main() {
	day, part, useExample := flags.ParseFlags()
	fmt.Println(useExample)
	switch day {
	case 1:
		day01.Day01(part, useExample)
	case 2:
		day02.Day02(part, useExample)
	}
}
