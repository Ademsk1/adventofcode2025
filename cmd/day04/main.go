package day04

import (
	printer "github.com/Ademsk1/adventofcode2025/util/print"
	"github.com/Ademsk1/adventofcode2025/util/read"
)

func Day04(part int, useExample bool) {
	content := read.Read("day04", useExample)
	data := read.DataAsGrid(content, "\n")
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
func part1(data [][]string) {
	sum := 0
	for i, row := range data {
		for j := range row {
			if data[i][j] == "@" {
				adjacentIndexes := getAdjacentIndexes(i, j, data)
				count := countRolls(data, adjacentIndexes)
				if count < 4 {
					sum++
				}
			}
		}
	}
	printer.PrettyPrintResult("3", "1", sum)
}

func countRolls(data [][]string, adjacentIndex [][][]int) int {
	count := 0
	for _, row := range adjacentIndex {
		for _, index := range row {
			if data[index[0]][index[1]] == "@" {
				count++
			}

		}
	}
	return count
}

func getAdjacentIndexes(i int, j int, data [][]string) [][][]int {
	// i acting as the index of the rows, j, being the index of the column
	var adjacentIndexes [][][]int = [][][]int{
		{
			{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1},
		},
		{
			{i, j - 1}, {i, j + 1},
		},
		{
			{i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1},
		},
	}
	counter := 0

	if j == 0 {
		for counter < 3 {
			adjacentIndexes[counter] = adjacentIndexes[counter][1:]
			counter++
		}
	} else if j == len(data[0])-1 {
		for counter < 3 {
			if counter == 1 {
				adjacentIndexes[counter] = adjacentIndexes[counter][:1]
			} else {
				adjacentIndexes[counter] = adjacentIndexes[counter][:2]
			}
			counter++
		}
	}
	if i == 0 {
		adjacentIndexes = adjacentIndexes[1:]
	} else if i == len(data)-1 {
		adjacentIndexes = adjacentIndexes[:2]
	}
	return adjacentIndexes
}
func part2(data [][]string) {
	var stored_positions [][]int
	rolls_removed := 0
	start := true
	for start || len(stored_positions) != 0 {
		start = false
		for _, indexes := range stored_positions {
			rolls_removed++
			data[indexes[0]][indexes[1]] = "."
		}
		stored_positions = [][]int{}
		for i, row := range data {
			for j := range row {
				if data[i][j] == "@" {
					adjacentIndexes := getAdjacentIndexes(i, j, data)
					count := countRolls(data, adjacentIndexes)
					if count < 4 {
						stored_positions = append(stored_positions, []int{i, j})
					}
				}
			}
		}
	}
	printer.PrettyPrintResult("4", "2", rolls_removed)
}
