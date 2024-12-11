package main

import (
	"fmt"
	"slices"
)

func run8(input []byte, step string) error {
	lines := splitNewLine(input)
	width, height, antennas := parseDay8(lines)

	if step == "a" {
		return run8a(width, height, antennas)
	}
	return run8b(width, height, antennas)
}

func parseDay8(lines []string) (width int, height int, antennas map[byte][]Vec2D) {
	antennas = make(map[byte][]Vec2D)
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				if antennas[byte(char)] == nil {
					antennas[byte(char)] = make([]Vec2D, 0)
				}
				antennas[byte(char)] = append(antennas[byte(char)], Vec2D{X: x, Y: y})
			}
		}
	}
	height = len(lines)
	width = len(lines[0])
	return
}

func run8a(width int, height int, antennas map[byte][]Vec2D) error {
	ans := make(map[Vec2D]bool)
	for _, values := range antennas {
		pairs := orderedPair(values)
		for _, pair := range pairs {
			an1 := Vec2D{
				X: 2*pair[0].X - pair[1].X,
				Y: 2*pair[0].Y - pair[1].Y,
			}
			an2 := Vec2D{
				X: 2*pair[1].X - pair[0].X,
				Y: 2*pair[1].Y - pair[0].Y,
			}
			if 0 <= an1.X && an1.X < width && 0 <= an1.Y && an1.Y < height {
				ans[an1] = true
			}
			if 0 <= an2.X && an2.X < width && 0 <= an2.Y && an2.Y < height {
				ans[an2] = true
			}
		}
	}
	// printAntennas(width, height, antennas, ans)
	fmt.Println(len(ans))
	return nil
}

func run8b(width int, height int, antennas map[byte][]Vec2D) error {
	ans := make(map[Vec2D]bool)
	for _, values := range antennas {
		pairs := orderedPair(values)
		for _, pair := range pairs {
			antinodeDir := Vec2D{
				X: pair[0].X - pair[1].X,
				Y: pair[0].Y - pair[1].Y,
			}
			for i := 0; i < 50; i++ {
				an := pair[0].Add(antinodeDir.Mult(i))
				if 0 <= an.X && an.X < width && 0 <= an.Y && an.Y < height {
					ans[an] = true
				} else {
					break
				}
			}
			antinodeDir = Vec2D{
				X: pair[1].X - pair[0].X,
				Y: pair[1].Y - pair[0].Y,
			}
			for i := 0; i < 50; i++ {
				an := pair[1].Add(antinodeDir.Mult(i))
				if 0 <= an.X && an.X < width && 0 <= an.Y && an.Y < height {
					ans[an] = true
				} else {
					break
				}
			}
		}
	}
	// printAntennas(width, height, antennas, ans)
	fmt.Println(len(ans))
	return nil
}

func printAntennas(width int, height int, antennas map[byte][]Vec2D, antinodes map[Vec2D]bool) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if v, ok := antinodes[Vec2D{X: x, Y: y}]; v && ok {
				fmt.Print("#")
				continue
			}
			f := false
			for symbol, xpos := range antennas {
				if slices.ContainsFunc(xpos, func(p Vec2D) bool { return p.X == x && p.Y == y }) {
					fmt.Print(string(symbol))
					f = true
					break
				}
			}
			if !f {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
