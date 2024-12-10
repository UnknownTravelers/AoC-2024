package main

import (
	"regexp"
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

func parseInts(str string, sep string) []int {
	split := strings.Split(str, sep)
	out := make([]int, len(split))
	for idx, s := range split {
		out[idx] = parseInt(s)
	}
	return out
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

// ParseRegex return a list of all match with
func ParseRegex(str string, reg *regexp.Regexp) []map[string]string {
	matches := reg.FindAllStringSubmatch(str, len(str)+1)
	result := make([]map[string]string, len(matches))
	for idx, match := range matches {
		result[idx] = make(map[string]string)
		for i, name := range reg.SubexpNames() {
			if i != 0 && name != "" {
				result[idx][name] = match[i]
			}
		}
	}
	return result
}
