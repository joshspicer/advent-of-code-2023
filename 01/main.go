package main

import (
	"base"
	"fmt"
	"strconv"
)

var digitsAsWords = []string{
	"one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

var digitsAsWordsReversed = []string{
	"eno", "owt", "eerht", "ruof",
	"evif", "xis", "neves", "thgie", "enin",
}

var digitMap = map[string]byte{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

var digitMapReversed = map[string]byte{
	"eno":   '1',
	"owt":   '2',
	"eerht": '3',
	"ruof":  '4',
	"evif":  '5',
	"xis":   '6',
	"neves": '7',
	"thgie": '8',
	"enin":  '9',
}

func findFirstNumber(line string, reversed bool, allowNumberAsWords bool) byte {
	arr := []byte(line)
	if reversed {
		arr = base.Reverse(arr)
	}
	for globalIdx, ch := range arr {
		// Simple case
		if ch >= '0' && ch <= '9' {
			return ch
		}

		peekAheadN := func(n int) byte {
			if globalIdx+n >= len(arr) {
				// Too far!
				return 0
			}
			return arr[globalIdx+n]
		}

		base.Debug("     Evaluating %c", ch)

		// Words case
		if allowNumberAsWords {
			corpus := digitsAsWords
			wordToByteMap := digitMap
			if reversed {
				// Reverse the word list
				corpus = digitsAsWordsReversed
				wordToByteMap = digitMapReversed
			}

			subIdx := 0
			matching := corpus

			for len(matching) > 0 {
				matching = base.Filter(matching, func(word string) bool {
					if subIdx >= len(word) {
						return false
					}
					return word[subIdx] == peekAheadN(subIdx)
				})
				base.Debug("     still matching: %v", matching)
				if len(matching) == 1 {
					match := matching[0]
					if subIdx+1 == len(match) {
						return wordToByteMap[match]
					}
					// If we drop out of here,
					// continue consuming string to find match that is of proper length
				}
				subIdx++
			}
			// If we drop out here, no word match was found at this position
		}

	}
	panic("Unexpected input")
}

func generic_run(input []string, allowNumbersAsWords bool) int {
	sum := 0

	for _, line := range input {
		base.Debug(">> %s", line)
		first := findFirstNumber(line, false, allowNumbersAsWords)
		base.Debug("  (!) first: %c", first)
		last := findFirstNumber(line, true, allowNumbersAsWords)
		base.Debug("  (!) last: %c", last)
		v, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			panic(err)
		}
		sum += v
	}

	fmt.Println(sum)
	return sum
}

func run1(input []string) int {
	return generic_run(input, false)
}

func run2(input []string) int {
	return generic_run(input, true)
}

func main() {
	run1(base.ReadExample1Lines())
	run2(base.ReadExample2Lines())
	run1(base.ReadInputLines())
	run2(base.ReadInputLines())
}
