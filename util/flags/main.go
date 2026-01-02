package flags

import "flag"

func ParseFlags() (int, int, bool) {
	dayPtr := flag.Int("day", 2, "Advent of code day")
	partPtr := flag.Int("part", 2, "advent of code part")
	exampleDataSetPtr := flag.Bool("example", true, "use example dataset")
	flag.Parse()
	day := *dayPtr
	part := *partPtr
	example := *exampleDataSetPtr

	return day, part, example

}
