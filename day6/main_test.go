package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCalculateFishLoop(t *testing.T) {
	fish := []int{3, 4, 3, 1, 2}
	actual := CalculateFishLoop(fish, 18)
	expected := []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("CalculateFishLoop() mismatch (-want +got):\n%s", diff)
	}
}
