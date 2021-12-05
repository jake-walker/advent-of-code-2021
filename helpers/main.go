package helpers

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput(name string) string {
	content, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("failed to read input file: %v", content)
	}
	return string(content)
}

func GetLines(content string) []string {
	return strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
}

func GetInputLines(name string) []string {
	content := GetInput(name)
	lines := GetLines(content)
	output := []string{}
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			output = append(output, line)
		}
	}
	return output
}

func SplitLines(content []string) [][]string {
	output := [][]string{}
	for i := 0; i < len(content); i++ {
		output = append(output, strings.Split(content[i], ""))
	}
	return output
}

func GetInputIntList(name string) []int {
	content := GetInputLines(name)
	output := make([]int, len(content))
	for i, str := range content {
		if str == "" {
			continue
		}
		j, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("failed to convert input file line to int: %v", err)
		}
		output[i] = j
	}
	return output
}

func StringSliceToIntSlice(in []string) []int {
	out := make([]int, len(in))
	for i, s := range in {
		out[i], _ = strconv.Atoi(s)
	}
	return out
}
