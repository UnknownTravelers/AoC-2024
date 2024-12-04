package main

import (
	"fmt"
	"regexp"
)

var regMulStrict = regexp.MustCompile(`mul\((?P<A>\d{1,3}),(?P<B>\d{1,3})\)`)
var regMulDoDont = regexp.MustCompile(`mul\((?P<A>\d{1,3}),(?P<B>\d{1,3})\)|(?P<do>do)\(\)|(?P<dont>don't)\(\)`)

func run3(input []byte, step string) error {
	if step == "a" {
		return run3a(input)
	}
	return run3b(input)
}

func run3a(input []byte) error {
	matches := ParseRegex(string(input), regMulStrict)
	sum := 0
	for _, match := range matches {
		a := parseInt(match["A"])
		b := parseInt(match["B"])
		sum += a * b
	}

	fmt.Println(sum)

	return nil
}

func run3b(input []byte) error {
	matches := ParseRegex(string(input), regMulDoDont)
	sum := 0
	do := true
	for _, match := range matches {
		if v := match["do"]; v != "" {
			do = true
		} else if v := match["dont"]; v != "" {
			do = false
		}
		aStr := match["A"]
		bStr := match["B"]
		if do && aStr != "" && bStr != "" {
			a := parseInt(aStr)
			b := parseInt(bStr)
			sum += a * b
		}
	}

	fmt.Println(sum)

	return nil
}
