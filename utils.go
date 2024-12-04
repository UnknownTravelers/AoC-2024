package main

import (
	"slices"
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

// splitNewLine and remove last empty line
func splitNewLine(input []byte) []string {
	lines := strings.Split(string(input), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func forEach[T any, O any](input []T, cb func(T) O) []O {
	out := make([]O, len(input))
	for k, v := range input {
		out[k] = cb(v)
	}
	return out
}

// slice returns a list with some elements removed
func slice[T any](input []T, start, n int) []T {
	input = slices.Clone(input)
	input = slices.Delete(input, start, start+n)
	return input
}

// Count every occurence of every value in the input list
func Count[T comparable](input []T) map[T]int {
	out := make(map[T]int)
	for _, v := range input {
		out[v]++
	}
	return out
}
