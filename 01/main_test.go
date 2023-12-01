package main

import (
	"base"
	"testing"
)

func TestExample(t *testing.T) {
	expected := 142
	result := run1(base.ReadExample1Lines())
	if result != expected {
		t.Errorf("ERR: %d != %d.", result, expected)
	}
}

func TestPart01(t *testing.T) {
	expected := 55123
	result := run1(base.ReadInputLines())
	if result != expected {
		t.Errorf("ERR: %d != %d.", result, expected)
	}
}

func TestPart02(t *testing.T) {
	expected := 55260
	result := run2(base.ReadInputLines())
	if result != expected {
		t.Errorf("ERR: %d != %d.", result, expected)
	}
}
