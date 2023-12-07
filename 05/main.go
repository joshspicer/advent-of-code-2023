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

func run1(input []string) int {
	seeds, almanac := parse(input)

	relationshipOrder := []Thing{Seed, Soil, Fertilizer, Water, Light, Temperature, Humidity, Location}

	lowestLocation := -1

	for _, seed := range seeds {
		tmp := seed
		for i := 0; i < len(relationshipOrder)-1; i++ {
			tmp = almanac.computeMappedLocation(tmp, relationshipOrder[i], relationshipOrder[i+1])
		}
		if lowestLocation == -1 || tmp < lowestLocation {
			lowestLocation = tmp
		}
	}

	fmt.Println(lowestLocation)
	return lowestLocation
}

func main() {
	// run1(base.ReadExample1Lines())
	run1(base.ReadInputLines())

	// run2(base.ReadExample1Lines())
	// run2(base.ReadInputLines())
}
