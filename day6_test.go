package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseMap(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  Board
	}{
		{
			name: "Small",
			input: []byte{
				'.', '.', '.', '\n',
				'^', '.', '#', '\n',
				'.', '#', '.', '\n',
				'.', '.', '.', '\n',
			},
			want: Board{
				tiles: [][]Tile{
					{Free, Free, Free},
					{Free, Free, Obstacle},
					{Free, Obstacle, Free},
					{Free, Free, Free},
				},
				playerPos:       Pos{X: 0, Y: 1},
				playerDirection: North,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseMap(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_Run(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want bool
	}{
		{
			name: "Loop",
			b: parseMap([]byte{
				'.', '#', '.', '.', '\n',
				'#', '^', '.', '#', '\n',
				'.', '.', '#', '.', '\n',
				'.', '.', '.', '.', '\n',
			}),
			want: true,
		},
		{
			name: "Not Loop",
			b: parseMap([]byte{
				'#', '.', '.', '\n',
				'^', '.', '#', '\n',
				'.', '#', '.', '\n',
				'.', '.', '.', '\n',
			}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.b.Run()
			require.Equal(t, tt.want, got)
		})
	}
}
