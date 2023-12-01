package base

import (
	"os"
	"strings"
)

func ReadPart01Lines() []string {
	return readLines("input.part01.txt")
}

func ReadPart02Lines() []string {
	return readLines("input.part02.txt")
}

func ReadExampleLines() []string {
	return readLines("input.example.txt")
}

func readLines(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}
