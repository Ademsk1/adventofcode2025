package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ademsk1/adventofcode2025/util/read"
)

func Day02(part int, useExample bool) {
	content := read.Read("day02", useExample)
	data := read.LinesAsArray(content, ",")

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

func checkRepeatsTwice(text string) bool {
	if len(text)%2 != 0 {
		return false
	}
	return text[:int(len(text)/2)] == text[int(len(text)/2):]
}

func part1(data []string) {

	var values []int64
	var sum int64
	for _, line := range data {
		lower, higher := getRanges(line)
		fmt.Println(lower, higher)
		starting, _ := strconv.ParseInt(lower, 10, 64)
		ending, _ := strconv.ParseInt(higher, 10, 64)
		for i := starting; i <= ending; i++ {
			text := fmt.Sprint(i)
			if checkRepeatsTwice(text) {
				values = append(values, i)
				sum += i
			}
		}
	}
	fmt.Println(values)
	fmt.Println(sum)
}

func CheckIsRepeating(text string) bool {

	if len(text) == 1 {
		return false
	}
	// edge case that cant be resolved by the 2nd for loop below
	if len(text) == 2 {
		return text[0] == text[1]
	}
	// numbers of odd lengths must have it so that each character is the same.

	for i := len(text) / 2; i >= 1; i-- {

		if len(text)%i != 0 {
			continue
		}
		// compare [j-> j+i] matches [j+i -> j+2*i], i.e. for 123123123, and i is 3, check 1st 123 with 2nd 123, and 2nd 123 with 3rd.
		var isMatching = true
		for j := 0; j < len(text)-i; j += i {
			if text[j:j+i] != text[j+i:j+2*i] {
				isMatching = false
				continue
			}
		}
		if isMatching {
			return true
		}
	}
	return false

}

func part2(data []string) {
	var values []int64
	var sum int64
	for _, line := range data {
		lower, higher := getRanges(line)
		fmt.Println(lower, higher)
		starting, _ := strconv.ParseInt(lower, 10, 64)
		ending, _ := strconv.ParseInt(higher, 10, 64)
		for i := starting; i <= ending; i++ {
			text := fmt.Sprint(i)
			// fmt.Println("Checking ", text)

			if CheckIsRepeating(text) {
				values = append(values, i)
				sum += i
			}
		}
	}
	fmt.Println(values)
	fmt.Println(sum)
}

func getRanges(line string) (string, string) {
	fullRange := strings.Split(line, "-")
	return fullRange[0], fullRange[1]
}
