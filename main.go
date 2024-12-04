// Package main ...
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Need day nummber + step, ex: 1 a, 15 b")
		return
	}
	day := os.Args[1]
	step := os.Args[2]

	go func() {
		if err := recover(); err != nil {
			fmt.Printf("panic %v \n", err)
		}
	}()
	run(day, step)
}

func run(day string, step string) error {
	file, err := os.Open(fmt.Sprintf("inputs/%v.input", day))
	if err != nil {
		return err
	}
	defer file.Close()

	bx, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	switch day {
	case "1":
		return run1(bx, step)
	}
	return nil
}
