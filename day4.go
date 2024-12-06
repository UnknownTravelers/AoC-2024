package main

import "fmt"

func run4(input []byte, step string) error {
	lines := splitNewLine(input)

	if step == "a" {
		return run4a(lines)
	}
	return run4b(lines)
}

func run4a(lines []string) error {
	sum := 0
	for row, line := range lines {
		for col := range line {
			sum += checkPos(lines, row, col)
		}
	}

	fmt.Println(sum)
	return nil
}

func run4b(lines []string) error {
	sum := 0
	for row, line := range lines {
		for col, char := range line {
			if char == 'A' && checkX(lines, row, col) {
				sum++
			}
		}
	}

	fmt.Println(sum)
	return nil
}

// ----- 4a -----

func checkPos(input []string, row int, col int) int {
	sum := 0
	if input[row][col] == 'X' {
		if checkNorth(input, row, col) {
			sum++
		}
		if checkSouth(input, row, col) {
			sum++
		}
		if checkEast(input, row, col) {
			sum++
		}
		if checkWest(input, row, col) {
			sum++
		}
		if checkNorthEast(input, row, col) {
			sum++
		}
		if checkNorthWest(input, row, col) {
			sum++
		}
		if checkSouthEast(input, row, col) {
			sum++
		}
		if checkSouthWest(input, row, col) {
			sum++
		}
	}
	return sum
}

func checkNorth(input []string, row, col int) bool {
	if row-3 < 0 {
		return false
	}
	return input[row-1][col] == 'M' && input[row-2][col] == 'A' && input[row-3][col] == 'S'
}

func checkSouth(input []string, row, col int) bool {
	if row+3 > len(input)-1 {
		return false
	}
	return input[row+1][col] == 'M' && input[row+2][col] == 'A' && input[row+3][col] == 'S'
}

func checkEast(input []string, row, col int) bool {
	if col+3 > len(input[row])-1 {
		return false
	}
	return input[row][col+1] == 'M' && input[row][col+2] == 'A' && input[row][col+3] == 'S'
}

func checkWest(input []string, row, col int) bool {
	if col-3 < 0 {
		return false
	}
	return input[row][col-1] == 'M' && input[row][col-2] == 'A' && input[row][col-3] == 'S'
}

func checkNorthEast(input []string, row, col int) bool {
	if row-3 < 0 || col+3 > len(input[row])-1 {
		return false
	}
	return input[row-1][col+1] == 'M' && input[row-2][col+2] == 'A' && input[row-3][col+3] == 'S'
}

func checkNorthWest(input []string, row, col int) bool {
	if row-3 < 0 || col-3 < 0 {
		return false
	}
	return input[row-1][col-1] == 'M' && input[row-2][col-2] == 'A' && input[row-3][col-3] == 'S'
}

func checkSouthEast(input []string, row, col int) bool {
	if row+3 > len(input)-1 || col+3 > len(input[row])-1 {
		return false
	}
	return input[row+1][col+1] == 'M' && input[row+2][col+2] == 'A' && input[row+3][col+3] == 'S'
}

func checkSouthWest(input []string, row, col int) bool {
	if row+3 > len(input)-1 || col-3 < 0 {
		return false
	}
	return input[row+1][col-1] == 'M' && input[row+2][col-2] == 'A' && input[row+3][col-3] == 'S'
}

// ----- 4b -----
func checkX(input []string, row, col int) bool {
	if row+1 > len(input)-1 || row-1 < 0 || col+1 > len(input[row])-1 || col-1 < 0 {
		return false
	}
	return (input[row-1][col-1] == 'M' && input[row+1][col+1] == 'S' || input[row-1][col-1] == 'S' && input[row+1][col+1] == 'M') && (input[row+1][col-1] == 'M' && input[row-1][col+1] == 'S' || input[row+1][col-1] == 'S' && input[row-1][col+1] == 'M')
}
