package main

import (
	"strconv"
	"strings"
)

func parseInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return v
}

func splitNewLine(input []byte) []string {
	lines := strings.Split(string(input), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

// Count every occurence of every value in the input list
func Count[T comparable](input []T) map[T]int {
	out := make(map[T]int)
	for _, v := range input {
		out[v]++
	}
	return out
}
