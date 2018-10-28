package main

import (
	"testing"
)

func TestKeys(t *testing.T) {
	keys := []string{
		"AKIAIOSFODNN7EXAMPLE",
		"testing",
	}
	m := matches(keys)
	if len(m) != 1 {
		t.Errorf("Expecting 1 match for call to matchKeys")
	}
}
