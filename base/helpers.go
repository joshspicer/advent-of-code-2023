package base

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsDebug() bool {
	return os.Getenv("ADVENT_DEBUG") == "true"
}

func Debug(format string, a ...interface{}) {
	if IsDebug() {
		fmt.Printf("[D] "+format+"\n", a...)
	}
}

func AtoiOrPanic(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
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

func ReadExample3Lines() []string {
	return readLines("example.3.txt")
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

func FilterWhitespace(data string) []string {
	return Filter(strings.Split(data, " "), func(s string) bool {
		// Remove whitespace or empty strings
		return s != "" && s != " "
	})
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

func Map[T, V any](data []T, f func(T) V) []V {

	output := make([]V, len(data))

	for i, x := range data {
		o := f(x)
		output[i] = o
	}

	return output
}

// From a single array, group elements into a list of arrays with size n
func GroupBy[T any](data []T, n int, requireAllGroupsEqualSize bool) [][]T {

	// Ensure each group is the same size
	if requireAllGroupsEqualSize && len(data)%n != 0 {
		panic(fmt.Sprintf("Cannot group %d elements into groups of %d", len(data), n))
	}

	output := make([][]T, 0)

	for i := 0; i < len(data); i += n {
		output = append(output, data[i:i+n])
	}

	return output
}

// =====
// GCD and LCM borrowed from
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// =====
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
