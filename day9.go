package main

import (
	"fmt"
	"slices"
)

// DiskA is a band of memory space for day9 step a
type DiskA struct {
	Blocs []int
}

type File struct {
	ID   int
	Size int
	Pos  int
}

// DiskB is a band of memory space for day9 step b
type DiskB struct {
	Files []*File
	MaxID int
}

func run9(input []byte, step string) error {
	if step == "a" {
		disk := parseDay9a(input)
		return run9a(disk)
	}
	disk := parseDay9b(input)
	return run9b(disk)
}

func parseDay9a(input []byte) *DiskA {
	disk := &DiskA{}
	vals := parseInts(string(input[:len(input)-1]), "")
	id := 0
	free := false
	for _, n := range vals {
		if free {
			disk.Blocs = append(disk.Blocs, ListN(n, -1)...)
		} else {
			disk.Blocs = append(disk.Blocs, ListN(n, id)...)
			id++
		}
		free = !free
	}
	return disk
}

func parseDay9b(input []byte) *DiskB {
	disk := &DiskB{}
	vals := parseInts(string(input[:len(input)-1]), "")
	id := 0
	free := false
	curPos := 0
	for _, n := range vals {
		if free {
			disk.Files = append(disk.Files, &File{
				ID:   -1,
				Size: n,
				Pos:  curPos,
			})
		} else {
			disk.Files = append(disk.Files, &File{
				ID:   id,
				Size: n,
				Pos:  curPos,
			})
			id++
		}
		curPos += n
		free = !free
	}
	disk.MaxID = id - 1
	return disk
}

func run9a(d *DiskA) error {
	for i := 0; i < len(d.Blocs); i++ {
		if d.Blocs[i] == -1 {
			d.Blocs[i] = d.PopLast()
		}
	}
	sum := 0
	for idx, v := range d.Blocs {
		if v != -1 {
			sum += idx * v
		}
	}

	fmt.Println(sum)

	return nil
}

func run9b(d *DiskB) error {
	d.Compress()
	sum := 0
	for _, file := range d.Files {
		sum += file.Checksum()
	}
	// 47871233927 < sum < 57900832028171
	fmt.Println(sum)

	return nil
}

// PopLast remove the last non-empty bloc of the memory and return its value
func (d *DiskA) PopLast() int {
	for i := len(d.Blocs) - 1; i >= 0; i-- {
		if d.Blocs[i] == -1 {
			continue
		}
		lastValue := d.Blocs[i]
		d.Blocs = d.Blocs[:i]
		return lastValue
	}
	return -1
}

// Compress the disk
func (d *DiskB) Compress() {
	for id := d.MaxID; id >= 0; id-- {
		file := d.GetFileByID(id)
		if file == nil {
			continue
		}

		pos := d.FindFreeSize(file.Size)
		if pos == -1 {
			continue
		}

		free := d.GetFileByPos(pos)
		if free == nil {
			continue
		}

		if free.Pos > file.Pos {
			continue
		}

		fileIndex := slices.Index(d.Files, file)
		d.Files = slices.Delete(d.Files, fileIndex, fileIndex+1)
		freeIndex := slices.Index(d.Files, free)
		if free.Size > file.Size {
			file.Pos = free.Pos
			free.Pos += file.Size
			free.Size -= file.Size

			d.Files = slices.Insert(d.Files, freeIndex, file)
		} else {
			file.Pos = free.Pos
			d.Files = slices.Delete(d.Files, freeIndex, freeIndex+1)
			d.Files = slices.Insert(d.Files, freeIndex, file)
		}
	}
}

// GetFileByID return a file by id
func (d *DiskB) GetFileByID(id int) *File {
	if id == -1 {
		return nil
	}
	for _, file := range d.Files {
		if file.ID == id {
			return file
		}
	}
	return nil
}

// GetFileByPos return a size by position
func (d *DiskB) GetFileByPos(pos int) *File {
	for _, file := range d.Files {
		if file.Pos == pos {
			return file
		}
	}
	return nil
}

// FindFreeSize return the first position with enough free space
func (d *DiskB) FindFreeSize(size int) int {
	for _, file := range d.Files {
		if file.ID == -1 && file.Size >= size {
			return file.Pos
		}
	}
	return -1
}

// Checksum return the checksum of the file
func (f File) Checksum() int {
	if f.ID == -1 {
		return 0
	}
	return f.ID * (2*f.Pos + f.Size - 1) * (f.Size) / 2
}

// ListN return a list of a defined size, filled with the given value
func ListN(size int, val int) []int {
	out := make([]int, size)
	for idx := range out {
		out[idx] = val
	}
	return out
}
