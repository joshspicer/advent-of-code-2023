package main

import (
	"base"
	"fmt"
	"strings"
)

func parseLine(line string) []int {
	split := strings.Split(line, " ")
	return base.Map(split, func(s string) int {
		return base.AtoiOrPanic(s)
	})
}

func differencesArray(ints []int) ([]int, bool) {
	arr := make([]int, len(ints)-1)
	allZeroes := true
	for i := 0; i < len(ints)-1; i++ {
		curr := ints[i]
		next := ints[i+1]
		v := next - curr
		if v != 0 {
			allZeroes = false
		}
		arr[i] = v
	}
	return arr, allZeroes
}

func run(input []string, part2 bool) int {
	// Map index of input to the (pyramid) arrays of differences
	differencesMap := map[int][][]int{}

	for idx, line := range input {
		ints := parseLine(line)
		for {
			differencesMap[idx] = append(differencesMap[idx], ints)
			arr, allZeroes := differencesArray(ints)
			ints = arr
			if allZeroes || len(ints) == 0 {
				differencesMap[idx] = append(differencesMap[idx], ints)
				break
			}
		}
		base.Debug(" %d =>  %v", idx, differencesMap[idx])
	}

	// Extrapolate
	result := 0
	for _, pyramid := range differencesMap {
		base.Debug("----")
		prev := 0
		final := 0
		for j := len(pyramid) - 1; j >= 0; j-- {
			currLineLength := len(pyramid[j])
			curr := pyramid[j][currLineLength-1]
			final = curr + prev
			debugSign := "+"
			if part2 {
				curr = pyramid[j][0]
				final = curr - prev
				debugSign = "-"
			}
			base.Debug("%d = %d %s %d", final, curr, debugSign, prev)
			prev = final
		}
		base.Debug("  %d", final)
		result += final

	}

	fmt.Println(result)
	return result
}

func run1(input []string) int {
	return run(input, false)
}

func run2(input []string) int {
	return run(input, true)
}

func main() {
	// run1(base.ReadExample1Lines())
	run1(base.ReadInputLines())
	// run2(base.ReadExample1Lines())
	run2(base.ReadInputLines())
}
