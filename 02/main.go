package main

import (
	"base"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	cubes map[string]int
}

var extractGameRegex = regexp.MustCompile(`^Game (\d+):(.*)$`)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseLine(line string) Game {
	game := Game{}

	// Parse out game ID
	matches := extractGameRegex.FindStringSubmatch(line)
	if matches == nil {
		panic("Invalid line. No game regex match.")
	}

	parsedId, err := strconv.Atoi(matches[1])
	if err != nil {
		panic("Failed to extract game id")
	}

	game.id = parsedId

	// Parse out rounds
	rounds := strings.Split(matches[2], ";")
	game.rounds = []Round{}
	for _, round := range rounds {
		var cubes map[string]int = map[string]int{}
		// Parse out cubes
		cubesParsed := strings.Split(round, ",")
		for _, cube := range cubesParsed {
			// Parse out cube color and count
			cubeParts := strings.Split(strings.TrimSpace(cube), " ")
			if len(cubeParts) != 2 {
				panic("Invalid cube format")
			}

			color := cubeParts[1]
			count, err := strconv.Atoi(cubeParts[0])
			if err != nil {
				panic("Failed to parse cube count")
			}
			cubes[color] = count

		}
		game.rounds = append(game.rounds, Round{cubes: cubes})
	}
	return game
}

func run1(input []string) int {
	games := []Game{}

	for _, line := range input {
		g := parseLine(line)
		games = append(games, g)
	}

	var roundConstraint Round = Round{
		cubes: map[string]int{"red": 12, "green": 13, "blue": 14},
	}

	sum := 0
	for _, g := range games {
		gamePossible := true
		for roundIdx, round := range g.rounds {
			if roundImpossible(roundConstraint, round) {
				base.Debug("Game %d is impossible (round=%d)", g.id, roundIdx)
				gamePossible = false
				break
			}
		}
		if gamePossible {
			sum += g.id
		}
	}
	fmt.Println(sum)
	return sum
}

func run2(input []string) int {
	games := []Game{}

	for _, line := range input {
		g := parseLine(line)
		games = append(games, g)
	}
	sum := 0
	for _, g := range games {
		base.Debug("Game %d", g.id)
		power := 1

		var minimumRequired Round = Round{
			cubes: map[string]int{},
		}

		// For each round update minimumRequired
		for _, round := range g.rounds {
			cubes := round.cubes

			for color, count := range cubes {
				if _, ok := minimumRequired.cubes[color]; !ok {
					minimumRequired.cubes[color] = count
				} else {
					minimumRequired.cubes[color] = max(minimumRequired.cubes[color], count)
				}
			}
		}

		for _, val := range minimumRequired.cubes {
			power *= val
		}

		base.Debug("   R=%d G=%d B=%d", minimumRequired.cubes["red"], minimumRequired.cubes["green"], minimumRequired.cubes["blue"])
		base.Debug("   Game %d power: %d", g.id, power)
		sum += power
	}

	fmt.Println(sum)
	return sum

}

func roundImpossible(constraint Round, current Round) bool {
	for color, max := range constraint.cubes {
		if current.cubes[color] > max {
			return true
		}
	}
	return false
}

func main() {
	// run2(base.ReadExample1Lines())
	run1(base.ReadInputLines())
	run2(base.ReadInputLines())
}
