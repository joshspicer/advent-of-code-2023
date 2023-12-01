package base

import (
	"fmt"
	"os"
	"strings"
)

func isDebug() bool {
	return os.Getenv("ADVENT_DEBUG") == "true"
}

func Debug(format string, a ...interface{}) {
	if isDebug() {
		fmt.Printf("[D] "+format+"\n", a...)
	}
}

func ReadInputLines() []string {
	return readLines("input.txt")
}

func ReadExample1Lines() []string {
	return readLines("example.1.txt")
}

func ReadExample2Lines() []string {
	return readLines("example.2.txt")
}

func readLines(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func Reverse[T any](list []T) []T {
	for i, j := 0, len(list)-1; i < j; {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return list
}

func Filter[T any](data []T, f func(T) bool) []T {

	acc := make([]T, 0, len(data))

	for _, e := range data {
		if f(e) {
			acc = append(acc, e)
		}
	}
	return acc
}
