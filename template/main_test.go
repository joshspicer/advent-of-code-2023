package main

import (
	"base"
	"testing"
)

func TestExample(t *testing.T) {
	expected := "TODO"
	result := run(base.ReadExampleLines())
	if result != expected {
		t.Errorf("ERR: %s != %s.", result, expected)
	}
}
