package main

import (
	"github.com/google/go-cmp/cmp"
	"github.com/jake-walker/advent-of-code-2021/helpers"
	"strings"
	"testing"
)

func TestLoadInput(t *testing.T) {
	input := helpers.GetLines("NNCB\n\n" +
		"CH -> B\n" +
		"HH -> N\n" +
		"CB -> H\n" +
		"NH -> C\n" +
		"HB -> C\n" +
		"HC -> B\n" +
		"HN -> C\n" +
		"NN -> C\n" +
		"BH -> H\n" +
		"NC -> B\n" +
		"NB -> B\n" +
		"BN -> B\n" +
		"BB -> N\n" +
		"BC -> B\n" +
		"CC -> N\n" +
		"CN -> C")
	actualTemplate, actualPairs := LoadInput(input)
	expectedTemplate := []string{"N", "N", "C", "B"}
	expectedPairs := map[string]string{
		"CH": "B",
		"HH": "N",
		"CB": "H",
		"NH": "C",
		"HB": "C",
		"HC": "B",
		"HN": "C",
		"NN": "C",
		"BH": "H",
		"NC": "B",
		"NB": "B",
		"BN": "B",
		"BB": "N",
		"BC": "B",
		"CC": "N",
		"CN": "C",
	}

	if diff := cmp.Diff(actualTemplate, expectedTemplate); diff != "" {
		t.Errorf("LoadInput() template = %v, want %v", actualTemplate, expectedTemplate)
	}

	if diff := cmp.Diff(actualPairs, expectedPairs); diff != "" {
		t.Errorf("LoadInput() pairs =\n%v", diff)
	}
}

func TestInsertPairs(t *testing.T) {
	pairs := map[string]string{
		"CH": "B",
		"HH": "N",
		"CB": "H",
		"NH": "C",
		"HB": "C",
		"HC": "B",
		"HN": "C",
		"NN": "C",
		"BH": "H",
		"NC": "B",
		"NB": "B",
		"BN": "B",
		"BB": "N",
		"BC": "B",
		"CC": "N",
		"CN": "C",
	}
	template := strings.Split("NBCCNBBBCBHCB", "")
	actual := InsertPairs(template, pairs)
	expected := strings.Split("NBBBCNCCNBBNBNBBCHBHHBCHB", "")

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("InsertPairs() =\n%v", diff)
	}
}

func TestInsertLoop(t *testing.T) {
	pairs := map[string]string{
		"CH": "B",
		"HH": "N",
		"CB": "H",
		"NH": "C",
		"HB": "C",
		"HC": "B",
		"HN": "C",
		"NN": "C",
		"BH": "H",
		"NC": "B",
		"NB": "B",
		"BN": "B",
		"BB": "N",
		"BC": "B",
		"CC": "N",
		"CN": "C",
	}
	template := []string{"N", "N", "C", "B"}
	actual := InsertLoop(template, pairs, 10)
	expected := 1588

	if actual != expected {
		t.Errorf("InsertLoop() = %v, want %v", actual, expected)
	}
}
