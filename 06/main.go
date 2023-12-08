package main

import (
	"base"
	"fmt"
	"regexp"
)

type Race struct {
	time     int // Time of the race
	distance int // Record distance
}

var numberRegex = regexp.MustCompile(`\d+`)

func parse(input []string) []Race {
	times := numberRegex.FindAllString(input[0], -1)
	distances := numberRegex.FindAllString(input[1], -1)
	races := make([]Race, len(times))
	for i := 0; i < len(times); i++ {
		race := Race{
			time:     base.AtoiOrPanic(times[i]),
			distance: base.AtoiOrPanic(distances[i]),
		}
		races[i] = race
	}
	return races
}

func (race Race) waysToWin() int {
	count := 0
	for chargeTime := 0; chargeTime < race.time; chargeTime++ {
		raceTimeRemaining := race.time - chargeTime
		speed := chargeTime
		distance := speed * raceTimeRemaining

		if distance > race.distance {
			count++
		}
	}
	return count
}

func run1(input []string) int {
	races := parse(input)
	base.Debug("%v", races)

	result := 1
	for _, race := range races {
		result *= race.waysToWin()
	}

	fmt.Println(result)
	return result
}

func run2() int {
	race := Race{
		time:     49979494,
		distance: 263153213781851,
	}

	result := race.waysToWin()
	fmt.Println(result)
	return result
}

func main() {
	// run1(base.ReadExample1Lines())
	run1(base.ReadInputLines())
	run2()
}
