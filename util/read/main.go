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
	} else if strings.HasSuffix(cwd, "adventofcode2025") {
		path = cwd + "/cmd/" + day
	} else {
		path = cwd
	}
	return path
}

func Read(day string, test bool) string {
	filepath := GetFilePath(day)
	fmt.Println(filepath)
	var file string = "example.txt"
	options, err := os.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}
	hasExampleFile := false
	hasDataFile := false
	fmt.Println("Files found are")
	fmt.Println(options)
	for _, option := range options {
		if option.Name() == "example.txt" {
			hasExampleFile = true
		}
		if option.Name() == "data.txt" {
			hasDataFile = true
		}
	}
	if !test && hasDataFile {
		file = "data.txt"
	}
	if !hasExampleFile {
		log.Fatal("No example.txt file found")
	}
	var fullFileWD = filepath + "/" + file

	data, err := os.ReadFile(fullFileWD)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func LinesAsArray(data string, sep string) []string {

	baseSplit := strings.Split(data, sep)
	for i := range baseSplit {
		baseSplit[i] = strings.TrimSpace(baseSplit[i])
	}
	return baseSplit
}
