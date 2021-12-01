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

func GetInputLines(name string) []string {
	content := GetInput(name)
	return strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
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
