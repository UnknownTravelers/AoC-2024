package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiskA_PopLast(t *testing.T) {
	tests := []struct {
		name  string
		d     *DiskA
		want  int
		want2 []int
	}{
		{
			name:  "OK",
			d:     &DiskA{Blocs: []int{0, 0, -1, 1, -1, 2}},
			want:  2,
			want2: []int{0, 0, -1, 1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.PopLast()
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.want2, tt.d.Blocs)
		})
	}
}

func Test_run9a(t *testing.T) {
	tests := []struct {
		name string
		d    *DiskA
		want []int
	}{
		{
			name: "OK",
			d:    &DiskA{Blocs: []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9}},
			want: []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := run9a(tt.d)
			require.NoError(t, err)
			require.Equal(t, tt.want, tt.d.Blocs)
		})
	}
}

func TestDiskB_Compress(t *testing.T) {
	tests := []struct {
		name string
		d    *DiskB
		want []*File
	}{
		{
			name: "A",
			d: &DiskB{
				Files: []*File{
					{
						ID:   0,
						Size: 2,
						Pos:  0,
					},
					{
						ID:   -1,
						Size: 3,
						Pos:  2,
					},
					{
						ID:   1,
						Size: 3,
						Pos:  5,
					},
				},
				MaxID: 1,
			},
			want: []*File{
				{
					ID:   0,
					Size: 2,
					Pos:  0,
				},
				{
					ID:   1,
					Size: 3,
					Pos:  2,
				},
			},
		},
		{
			name: "B",
			d: &DiskB{
				Files: []*File{
					{
						ID:   0,
						Size: 2,
						Pos:  0,
					},
					{
						ID:   -1,
						Size: 4,
						Pos:  2,
					},
					{
						ID:   1,
						Size: 3,
						Pos:  6,
					},
					{
						ID:   -1,
						Size: 2,
						Pos:  9,
					},
					{
						ID:   2,
						Size: 3,
						Pos:  11,
					},
				},
				MaxID: 2,
			},
			want: []*File{
				{
					ID:   0,
					Size: 2,
					Pos:  0,
				},
				{
					ID:   2,
					Size: 3,
					Pos:  2,
				},
				{
					ID:   -1,
					Size: 1,
					Pos:  5,
				},
				{
					ID:   1,
					Size: 3,
					Pos:  6,
				},
				{
					ID:   -1,
					Size: 2,
					Pos:  9,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Compress()
			require.Equal(t, tt.want, tt.d.Files)
		})
	}
}

func TestFile_Checksum(t *testing.T) {
	tests := []struct {
		name string
		d    *File
		want int
	}{
		{
			name: "ID 0",
			d: &File{
				ID:   0,
				Size: 3,
				Pos:  2,
			},
			want: 0,
		},
		{
			name: "Empty Zone",
			d: &File{
				ID:   -1,
				Size: 3,
				Pos:  2,
			},
			want: 0,
		},
		{
			name: "Big",
			d: &File{
				ID:   50,
				Size: 20,
				Pos:  200,
			},
			want: 209500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.Checksum()
			require.Equal(t, tt.want, got)
		})
	}
}
