package main

import "testing"

func TestSum(t *testing.T) {
	actual := sum_of(2, 3)
	expected := 5

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
