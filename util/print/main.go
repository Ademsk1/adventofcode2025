package printer

import "fmt"

func PrettyPrintResult(day string, part string, result any) {

	fmt.Println("-------")
	fmt.Println("Day ", day, " | Part ", part )
	fmt.Println(result)
	fmt.Println("-------")
}