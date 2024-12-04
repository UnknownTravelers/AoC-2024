package main

import (
	"fmt"
	"sort"
	"strings"
)

func run1(input []byte, step string) error {
	lines := splitNewLine(input)

	lefts := make([]int, len(lines))
	rights := make([]int, len(lines))
	for idx, line := range lines {
		vals := strings.Split(line, "   ")
		lefts[idx] = parseInt(vals[0])
		rights[idx] = parseInt(vals[1])
	}

	if step == "a" {
		return run1a(lefts, rights)
	}
	return run1b(lefts, rights)
}

func run1a(lefts []int, rights []int) error {
	sort.Sort(sort.IntSlice(lefts))
	sort.Sort(sort.IntSlice(rights))

	sum := 0
	for idx := range lefts {
		dif := lefts[idx] - rights[idx]
		if dif > 0 {
			sum += dif
		} else {
			sum -= dif
		}
	}

	fmt.Println(sum)

	return nil
}

func run1b(lefts []int, rights []int) error {
	left := Count(lefts)
	right := Count(rights)

	sum := 0
	for val, leftCount := range left {
		rightCount, _ := right[val]
		sum += val * leftCount * rightCount
	}

	fmt.Println(sum)

	return nil
}
