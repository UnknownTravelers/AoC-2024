package main

import (
	"fmt"
	"strings"
)

func run2(input []byte, step string) error {
	lines := splitNewLine(input)

	if step == "a" {
		v, err := run2a(lines)
		fmt.Println(v)
		return err
	}
	v, err := run2b(lines)
	fmt.Println(v)
	return err
}

func run2a(reports []string) (int, error) {
	safeReports := 0
	for _, report := range reports {
		pointsStr := strings.Split(report, " ")
		levels := forEach(pointsStr, parseInt)
		if safeReport(levels) {
			safeReports++
		}
	}

	return safeReports, nil
}

func run2b(reports []string) (int, error) {
	safeReports := 0
	for _, report := range reports {
		pointsStr := strings.Split(report, " ")
		levels := forEach(pointsStr, parseInt)
		if safeReport(levels) {
			safeReports++
		} else {
			for idx := 0; idx < len(levels); idx++ {
				partialLevels := slice(levels, idx, 1)
				if safeReport(partialLevels) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports, nil
}

func safeLevels(growth bool, a, b int) bool {
	dif := b - a
	if dif > 3 || dif < -3 || dif == 0 {
		return false
	}
	return growth == (dif > 0)
}

func safeReport(levels []int) bool {
	growth := levels[0] < levels[1] // growth direction | increase = true, decrease = false
	for ptr := 0; ptr < len(levels)-1; ptr++ {
		if !safeLevels(growth, levels[ptr], levels[ptr+1]) {
			return false
		}
	}
	return true
}

func safeReportWithDampener(levels []int) bool {
	pdUsed := false
	growth := levels[0] < levels[1] // growth direction | increase = true, decrease = false
	for ptr := 0; ptr < len(levels)-1; ptr++ {
		if !safeLevels(growth, levels[ptr], levels[ptr+1]) {
			if pdUsed {
				return false
			}
			if ptr >= len(levels)-2 { // removing last value will assure this report is safe
				return true
			}
			if !safeLevels(growth, levels[ptr], levels[ptr+2]) {
				return false
			}
		}
	}
	return true
}
