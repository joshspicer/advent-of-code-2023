package main

import (
	"base"
	"testing"
)

func TestExample(t *testing.T) {
	expected := 142
	result := run(base.ReadExampleLines())
	if result != expected {
		t.Errorf("ERR: %d != %d.", result, expected)
	}
}
