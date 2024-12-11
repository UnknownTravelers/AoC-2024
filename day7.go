package main

import (
	"fmt"
	"regexp"
	"slices"
)

var regEq = regexp.MustCompile(`^(?P<total>\d+): (?P<components>\d+(?: \d+?)*)$`)

// Op operation
type Op func(int, int) int

type Equation struct {
	Total     int
	Component []int
}

func run7(input []byte, step string) error {
	lines := splitNewLine(input)
	eqx := make([]Equation, len(lines))
	for idx, line := range lines {
		m := ParseRegex(line, regEq)
		eqx[idx] = Equation{
			Total:     parseInt(m[0]["total"]),
			Component: parseInts(m[0]["components"], " "),
		}
	}
	if step == "a" {
		return run7a(eqx)
	}
	return run7b(eqx)
}

func run7a(eqx []Equation) error {
	validOps := []Op{opAdd, opMult}
	sum := 0
	for _, eq := range eqx {
		tmp := eq.Component[0]
		results := tryOp(tmp, eq.Component[1:], validOps)
		if slices.Contains(results, eq.Total) {
			sum += eq.Total
		}
	}

	fmt.Println(sum)
	return nil
}

func run7b(eqx []Equation) error {
	validOps := []Op{opAdd, opMult, opConcat}
	sum := 0
	for _, eq := range eqx {
		tmp := eq.Component[0]
		results := tryOp(tmp, eq.Component[1:], validOps)
		if slices.Contains(results, eq.Total) {
			sum += eq.Total
		}
	}

	fmt.Println(sum)
	return nil
}

func opMult(a int, b int) int {
	return a * b
}

func opAdd(a int, b int) int {
	return a + b
}

func opConcat(a int, b int) int {
	return parseInt(fmt.Sprintf("%v%v", a, b))
}

func tryOp(curValue int, values []int, ops []Op) []int {
	if len(values) == 0 {
		return []int{curValue}
	}
	results := make([]int, 0)
	for _, op := range ops {
		tmpVal := op(curValue, values[0])
		results = append(results, tryOp(tmpVal, values[1:], ops)...)
	}
	return results
}
