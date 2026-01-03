package day03

import (
	"fmt"
	"log"
	"math"
	"strconv"

	printer "github.com/Ademsk1/adventofcode2025/util/print"
	"github.com/Ademsk1/adventofcode2025/util/read"
)

func Day03(part int, useExample bool) {
	content := read.Read("day03", useExample)
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

func part1(data []string) {
	fmt.Println(data)
	// employ sliding window technique
	var sum int64 = 0
	for _, line := range data {
		joltage := getHighestJoltage(line)
		sum += joltage
	}
	printer.PrettyPrintResult("3", "1", sum)

}

func convertToJoltage(char1 byte, char2 byte) int64 {
	joltage, err := strconv.ParseInt(string(char1)+string(char2), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return joltage
}

func convertToArray(s string) []int64 {
	var arr []int64
	for _, char := range s {

		v, err := strconv.ParseInt(string(char), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, v)
	}
	return arr
}

func getMaxIndex(arr []int64, subtractor int) int {
	var currentMaxIndex int = 0
	// len(arr) - 1 because we can't accept the last value as the first char to be used

	for i := 0; i < len(arr)-subtractor; i++ {
		if arr[i] > arr[currentMaxIndex] {
			currentMaxIndex = i
		}
	}
	return currentMaxIndex
}

func getHighestJoltage(line string) int64 {
	joltages := convertToArray(line)
	fmt.Println(joltages)
	maxIndex := getMaxIndex(joltages, 1)
	fmt.Println("Max ", string(line[maxIndex]))

	secondMax := getMaxIndex(joltages[maxIndex+1:], 0)
	secondMax += maxIndex + 1
	fmt.Println("2nd Max ", string(line[secondMax]))
	v := convertToJoltage(line[maxIndex], line[secondMax])
	fmt.Println("Value ", v)

	return v

}

func part2(data []string) {
	var sum int64
	for _, line := range data {
		sum += get12BatteryJoltages(line)
	}
	printer.PrettyPrintResult("3", "2", sum)

}

func get12BatteryJoltages(batteries string) int64 {
	joltages := convertToArray(batteries)
	var joltageVal float64 = 0
	for subtractor := 11; subtractor >= 0; subtractor-- {
		maxIndex := getMaxIndex(joltages, subtractor)
		joltageVal += math.Pow(10, float64(subtractor)) * float64(joltages[maxIndex])
		joltages = joltages[maxIndex+1:]

	}
	fmt.Println(int(joltageVal))
	return int64(joltageVal)
}
