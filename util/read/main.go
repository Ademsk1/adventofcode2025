package read

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetFilePath(day string) string {
	var path string
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cwd)
	if strings.HasSuffix(cwd, "cmd") {
		path = cwd + "/" + day
	} else {
		path = cwd
	}
	return path
}

func Read(day string, test bool) string {
	filepath := GetFilePath(day)
	var file string
	if test {
		file = "example.txt"
	} else {
		file = "data.txt"
	}
	var fullFileWD = filepath + "/" + file
	data,err := os.ReadFile(fullFileWD)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func LinesAsArray(data string) []string {

	baseSplit := strings.Split(data, "\n")
	for i, _ := range baseSplit {
		strings.TrimSpace(baseSplit[i])
	}
	return baseSplit
}