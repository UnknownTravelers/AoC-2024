package main

import (
	"fmt"
)

// Tile type in a board
type Tile int

const (
	// Free space
	Free Tile = iota
	// Obstacle blocked space
	Obstacle
)

// Direction is a Cardinal direction
type Direction int

const (
	// North direction
	North Direction = iota
	//East direction
	East
	// South direction
	South
	// West direction
	West
)

// Pos in a 2D plane
type Pos struct {
	X, Y int
}

// Board Tile 2d plane
type Board struct {
	tiles           [][]Tile
	playerPos       Pos
	playerDirection Direction
}

func run6(input []byte, step string) error {
	board := parseMap(input)

	if step == "a" {
		return run6a(board)
	}
	return run6b(board)
}

func run6a(board Board) error {
	visitedTiles := map[Pos]bool{board.playerPos: true}
	for {
		if board.Step() {
			break
		}
		visitedTiles[board.playerPos] = true

	}
	sum := 0
	for _, v := range visitedTiles {
		if v {
			sum++
		}
	}
	fmt.Println(sum)

	return nil
}

func run6b(board Board) error {
	sum := 0
	originalPos := board.playerPos
	visitedTiles := map[Pos]bool{originalPos: true}
	for {
		if board.Step() {
			break
		}
		visitedTiles[board.playerPos] = true
	}

	for k, v := range visitedTiles {
		if !v {
			fmt.Println("skip: tile was not visited")
			continue
		}
		// skip original player pos
		if k == originalPos {
			fmt.Println("skip: original pos")
			continue
		}

		tiles := make([][]Tile, len(board.tiles))
		for y, line := range board.tiles {
			tiles[y] = make([]Tile, len(line))
			for x, tile := range line {
				tiles[y][x] = tile
			}
		}

		b := Board{
			tiles:           tiles,
			playerPos:       originalPos,
			playerDirection: North,
		}

		b.tiles[k.Y][k.X] = Obstacle
		if b.Run() {
			sum++
		}
	}

	fmt.Println(sum)
	return nil
}

func parseMap(input []byte) Board {
	board := Board{}
	tiles := make([][]Tile, 0)

	curRow := make([]Tile, 0)
	for _, c := range input {
		switch c {
		case '\n':
			tiles = append(tiles, curRow)
			curRow = make([]Tile, 0)
		case '.':
			curRow = append(curRow, Free)
		case '#':
			curRow = append(curRow, Obstacle)
		case '^':
			curRow = append(curRow, Free)
			board.playerDirection = North
			board.playerPos = Pos{Y: len(tiles), X: len(curRow) - 1}
		}
	}
	board.tiles = tiles
	return board
}

// Step make player advance or rotate 90° to the right. If player leave map, return true
func (b *Board) Step() bool {
	switch b.playerDirection {
	case North:
		next := Pos{b.playerPos.X, b.playerPos.Y - 1}
		if next.Y < 0 {
			return true
		}
		if b.tiles[next.Y][next.X] == Free {
			b.playerPos = next
		} else {
			b.playerDirection = East
		}
	case East:
		next := Pos{b.playerPos.X + 1, b.playerPos.Y}
		if next.X >= len(b.tiles[next.Y]) {
			return true
		}
		if b.tiles[next.Y][next.X] == Free {
			b.playerPos = next
		} else {
			b.playerDirection = South
		}
	case South:
		next := Pos{b.playerPos.X, b.playerPos.Y + 1}
		if next.Y >= len(b.tiles) {
			return true
		}
		if b.tiles[next.Y][next.X] == Free {
			b.playerPos = next
		} else {
			b.playerDirection = West
		}
	case West:
		next := Pos{b.playerPos.X - 1, b.playerPos.Y}
		if next.X < 0 {
			return true
		}
		if b.tiles[next.Y][next.X] == Free {
			b.playerPos = next
		} else {
			b.playerDirection = North
		}
	}
	return false
}

// Run the board position. If stuck in a loop, return true
func (b *Board) Run() bool {
	type CompletePos struct {
		Pos Pos
		Dir Direction
	}
	visitedTiles := map[CompletePos]bool{}
	for {
		if b.Step() {
			return false
		}
		if v, ok := visitedTiles[CompletePos{Pos: b.playerPos, Dir: b.playerDirection}]; v && ok {
			return true
		}
		visitedTiles[CompletePos{Pos: b.playerPos, Dir: b.playerDirection}] = true
	}
}

// Print display board with player path
func (b *Board) Print(visitedTiles map[Pos]bool) {
	for i, line := range b.tiles {
		for j, tile := range line {
			if tile == Free {
				if v, ok := visitedTiles[Pos{X: j, Y: i}]; ok && v {
					fmt.Print("X")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}
