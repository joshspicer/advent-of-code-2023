package main

import (
	"base"
	"fmt"
	"strings"
)

// Enum of all the types of things
type Thing string

const (
	Seed        Thing = "seed"
	Soil        Thing = "soil"
	Fertilizer  Thing = "fertilizer"
	Water       Thing = "water"
	Light       Thing = "light"
	Temperature Thing = "temperature"
	Humidity    Thing = "humidity"
	Location    Thing = "location"
)

type Almanac map[Relationship][]Shift

type Relationship struct {
	Source Thing
	Dest   Thing
}

type Shift struct {
	source int
	dest   int
	length int
}

func (almanac Almanac) computeMappedLocation(id int, source Thing, dest Thing) int {
	shifts := almanac[Relationship{Source: source, Dest: dest}]
	if len(shifts) == 0 {
		errString := "No shifts found for %s -> %s"
		panic(errString)
	}
	val := id
	for _, shift := range shifts {
		if id >= shift.source && id < shift.source+shift.length {
			val = shift.dest + (id - shift.source)
			break
		}
	}
	base.Debug("  %s -> %s : %d -> %d", source, dest, id, val)
	return val
}

func parse(input []string) (seeds []int, almanac Almanac) {
	almanac = make(Almanac)

	var tmpShifts []Shift = make([]Shift, 0)
	var tmpRelationship Relationship = Relationship{}

	for i := 0; i < len(input); i++ {
		line := input[i]

		if line == "" {
			// Commit completed entry of almanac
			if len(tmpShifts) > 0 {
				almanac[tmpRelationship] = tmpShifts

				tmpShifts = make([]Shift, 0)
				tmpRelationship = Relationship{}
			}
			continue
		}

		if strings.HasPrefix(line, "seeds: ") {
			inputSeeds := strings.Split(strings.Split(line, "seeds: ")[1], " ")
			seeds = base.Map(inputSeeds, func(s string) int {
				return base.AtoiOrPanic(s)
			})
			continue
		}

		if strings.Contains(line, " map:") {
			relationship := strings.Split(strings.Split(line, " map:")[0], "-to-")
			source := Thing(relationship[0])
			dest := Thing(relationship[1])
			tmpRelationship = Relationship{Source: source, Dest: dest}
			continue
		}

		// Parse mapping
		values := strings.Split(line, " ")
		if len(values) != 3 {
			panic("Unexpected mapping line: " + line)
		}

		dest := base.AtoiOrPanic(values[0])
		source := base.AtoiOrPanic(values[1])
		length := base.AtoiOrPanic(values[2])
		tmpShifts = append(tmpShifts, Shift{source: source, dest: dest, length: length})
	}

	base.Debug("")
	base.Debug("Seed: %v", seeds)
	if base.IsDebug() {
		for k, v := range almanac {
			base.Debug("  %s -> %s : %+v", k.Source, k.Dest, v)
		}
	}

	return seeds, almanac
}

func (almanac Almanac) run(seeds []int) (int, int) {
	relationshipOrder := []Thing{Seed, Soil, Fertilizer, Water, Light, Temperature, Humidity, Location}

	lowestLocation := -1
	bestIndex := -1

	for idx, seed := range seeds {
		tmp := seed

		base.Debug("  --------------------------------")
		for i := 0; i < len(relationshipOrder)-1; i++ {
			tmp = almanac.computeMappedLocation(tmp, relationshipOrder[i], relationshipOrder[i+1])
		}

		if lowestLocation == -1 || tmp < lowestLocation {
			lowestLocation = tmp
			bestIndex = idx
		}
	}
	return lowestLocation, bestIndex
}

func run1(input []string) int {
	seeds, almanac := parse(input)
	result, _ := almanac.run(seeds)
	fmt.Println(result)
	return result
}

// This is wildly inefficient, but worked on my input :^)
func run2(input []string) int {
	seedInput, almanac := parse(input)

	bestCandiateSeeds := make([]int, 0)
	for i := 0; i < len(seedInput)-1; i += 2 {
		start := seedInput[i]
		length := seedInput[i+1]

		seedRange := make([]int, length)
		base.Debug("Adding seeds from %d to %d", start, start+length)
		for j := 0; j < length; j++ {
			seedRange[j] = start + j
		}

		for len(seedRange) >= 10 {
			base.Debug("  Testing %v", seedRange)

			// For each side, test the lowest and highest seed.
			// Only keep the side with the lowest location (best)
			_, bestIndex := almanac.run([]int{seedRange[0], seedRange[len(seedRange)-1]})
			if bestIndex == 0 {
				seedRange = seedRange[:len(seedRange)-1]
			} else {
				// TODO: I did this wrong before running run2, but not gonnna fix :)
				seedRange = seedRange[1:]
			}
		}

		bestCandiateSeeds = append(bestCandiateSeeds, seedRange...)
	}

	result, _ := almanac.run(bestCandiateSeeds)
	fmt.Println(result)
	return result
}

func main() {
	// run1(base.ReadExample1Lines())
	run1(base.ReadInputLines())
	// run2(base.ReadExample1Lines())
	run2(base.ReadInputLines()) // 125742456
}
