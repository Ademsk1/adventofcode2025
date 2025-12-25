package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Ademsk1/adventofcode2025/util/read"
)

func main() {
	var usingExample = false
	content := read.Read("day01", usingExample)
	data := read.LinesAsArray(content)
	part := flag.Int("part", 2, "advent of code part (1/2)")
	flag.Parse()
	fmt.Println(*part)
	if *part == 1 {
		part1(data)
	} else if *part == 2 {
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
	fmt.Println("-----")

	fmt.Println("Starting part 1...")
	for _, line := range data {
		if line == "" {
			continue
		}
		ParseTurn(line, &rotationSize, &sign)

		dial += (int(rotationSize) * sign)

		if (dial % 100 == 0) {
			zeroCounter++
		}
		
	}
	fmt.Println("Part 1: Zero counter was ", zeroCounter)
	fmt.Println("-----")

}

func part2(data []string) {
	// zeroCounter := 0
	// dial := 50
	// sign := -1
	// for _, line := range data {
		
	// }


}