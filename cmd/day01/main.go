package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	printer "github.com/Ademsk1/adventofcode2025/util/print"
	"github.com/Ademsk1/adventofcode2025/util/read"
)

func main() {

	partPtr := flag.Int("part", 2, "advent of code part")
	exampleDataSetPtr := flag.Bool("example", true, "use example dataset")

	flag.Parse()
	var part int = *partPtr
	content := read.Read("day01", *exampleDataSetPtr)
	data := read.LinesAsArray(content)
	if part == 1 {
		part1(data)
	} else if part == 2 {
		comparemethods(data)
	} else {
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

		if (dial % 100 == 0) {
			zeroCounter++
		}
		
	}
	printer.PrettyPrintResult("1", "1", zeroCounter)


}

func part2(data []string) {
	zeroCounter := 0
	dial := 50
	newDial := 0
	sign := -1
	rotationSize := 0
	for _, line := range data {
		if line == "" {
			continue 
		}
		ParseTurn(line, &rotationSize, &sign)

		newDial = dial + (rotationSize * sign)
		if dial > 0 && newDial <=0 {
			zeroCounter += int(math.Floor(float64(-newDial)/100)) + 1
			dial = 100 + (newDial % 100)
		} else if newDial >= 100 {
			zeroCounter += int(math.Floor(float64(newDial)/100))
			dial = newDial % 100
		} else {
			dial = newDial %100
		}
		// } else if newDial == 0 {
		// 	zeroCounter++
		// }
		
	}
	printer.PrettyPrintResult("1", "2", zeroCounter)
}


func comparemethods(data []string) {
	dial1 := 50
	dial2 := 50

	var rotationSize int
	var sign int

	for _, line := range data {
		if line == "" {
			continue
		}
		ParseTurn(line, &rotationSize, &sign)
		d1, z1 := alterDial(rotationSize, sign, dial1)
		d2, z2 := slowpart2(rotationSize, sign, dial2)
		if (d1 != d2) {
			fmt.Println("Mismatch between slow and fast methods : ")
			fmt.Printf("New dial\n")
			fmt.Printf("Fast : %v , Slow : %v\n", d1, d2)
			fmt.Println("Starting at : ")
			fmt.Printf("Fast : %v, Slow : %v\n", dial1, dial2)
			panic(1)
		}
		dial1 = d1
		dial2 = d2
		if (z1 != z2) {
			fmt.Println("Mismatch between slow and fast zero counter : ")
			fmt.Printf("New dial\n")
			fmt.Printf("Fast : %v , Slow : %v\n", d1, d2)
			fmt.Println("Starting at : ")
			fmt.Printf("Fast : %v, Slow : %v\n", dial1, dial2)
			fmt.Println("Zero counter : ")
			fmt.Printf("Fast : %v, Slow : %v\n", z1, z2)
			panic(1)
		}


	}
}



func alterDial(rotationSize int, sign int, dial int) (int, int) {
	var zeroCounter int = 0
	newDial := dial + (rotationSize * sign)
	leftBound := 0
	rightBound := 100
	if (dial == 0) {
		leftBound = -100
	}
	
	if dial > 0 && newDial <=0 {
		zeroCounter += int(math.Floor(float64(-newDial)/100)) + 1
		dial = 100 + (newDial % 100)
	} else if newDial >= 100 {
		zeroCounter += int(math.Floor(float64(newDial)/100))
		dial = newDial % 100
	} else {
		dial = newDial %100
	}
		// } else if newDial == 0 {
		// 	zeroCounter++
		// }
	return dial, zeroCounter
	
}


func slowpart2(rotationSize int, sign int, dial int) (int, int) {
	var zeroCounter int = 0
	for i := 0; i < rotationSize; i ++ {
			dial += sign
			if dial == 0 {
				zeroCounter++
				
			}
			if dial < 0 {
				dial = 100 + dial
			}
			if dial == 100 {
				dial = 0
				zeroCounter++
			}


		}

	return dial, zeroCounter
}