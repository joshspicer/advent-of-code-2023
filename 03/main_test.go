package main

import (
	"base"
	"testing"
)

func Test(t *testing.T) {
	part1 := 527369
	part2 := 73074886
	actual1 := run1(base.ReadInputLines())
	actual2 := run2(base.ReadInputLines())
	if part1 != actual1 {
		t.Errorf("ERR: %d != %d.", part1, actual1)
	}
	if part2 != actual2 {
		t.Errorf("ERR: %d != %d.", part2, actual2)
	}
}
