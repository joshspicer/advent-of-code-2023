package main

import (
	"base"
	"testing"
)

func TestExample(t *testing.T) {
	expected := "TODO"
	result := run1(base.ReadExample1Lines())
	if result != expected {
		t.Errorf("ERR: %s != %s.", result, expected)
	}
}
