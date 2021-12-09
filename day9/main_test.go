package main

import (
	"github.com/google/go-cmp/cmp"
	"sort"
	"testing"
)

var input = "2199943210\n3987894921\n9856789892\n8767896789\n9899965678"

func TestLoadInput(t *testing.T) {
	actual := LoadInput(input)
	expected := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("LoadInput() =\n%v", diff)
	}
}

func TestGetLowPoints(t *testing.T) {
	heightmap := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	actual := GetLowPoints(heightmap)
	expected := []int{1, 0, 5, 5}

	sort.Ints(actual)
	sort.Ints(expected)

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("GetLowPoints() =\n%v", diff)
	}
}
