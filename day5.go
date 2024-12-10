package main

import (
	"fmt"
	"slices"
	"strings"
)

type Rules map[int][]int

type Document struct {
	Pages []int
}

type Pair [2]int

func run5(input []byte, step string) error {
	lines := splitNewLine(input)
	rules, docs := parseDay5(lines)

	if step == "a" {
		return run5a(rules, docs)
	}
	return run5b(rules, docs)
}

func run5a(rules Rules, docs []Document) error {
	sum := 0
	for _, doc := range docs {
		if doc.CheckRules(rules) {
			sum += doc.Pages[(len(doc.Pages)-1)/2]
		}
	}

	fmt.Println(sum)

	return nil
}

func run5b(rules Rules, docs []Document) error {
	sum := 0
	for _, doc := range docs {
		if !doc.CheckRules(rules) {
			doc.Sort(rules)
			sum += doc.Pages[(len(doc.Pages)-1)/2]
		}
	}

	// 3974 < sum
	fmt.Println(sum)

	return nil
}

func parseDay5(lines []string) (Rules, []Document) {
	rules := make(Rules)
	curLine := 0
	for idx, line := range lines {
		if line == "" {
			curLine = idx + 1
			break
		}
		nbs := strings.Split(line, "|")
		prev := parseInt(nbs[0])
		next := parseInt(nbs[1])
		if rules[prev] == nil {
			rules[prev] = make([]int, 0)
		}
		rules[prev] = append(rules[prev], next)
	}

	documents := make([]Document, 0)
	for _, line := range lines[curLine:] {
		documents = append(documents, Document{
			Pages: parseInts(line, ","),
		})
	}
	return rules, documents
}

// CheckRules check if a document follows given rules
func (d *Document) CheckRules(rules Rules) bool {
	for _, p := range orderedPair(d.Pages) {
		if !Pair(p).CheckRules(rules) {
			return false
		}
	}
	return true
}

// Sort sort document pages with the given rules
func (d *Document) Sort(rules Rules) {
	sorted := false
	for !sorted {
		sorted = true
		for _, p := range orderedPair(d.Pages) {
			if !Pair(p).CheckRules(rules) {
				sorted = false
				// move p2 to just before p1
				d.MovePage(Pair(p))
				break
			}
		}
	}
}

// MovePage move p2 to just before p1
func (d *Document) MovePage(pair Pair) {
	p1 := slices.Index(d.Pages, pair[0])
	p2 := slices.Index(d.Pages, pair[1])
	pages := slice(d.Pages, p2, 1)
	d.Pages = slices.Insert(pages, p1, pair[1])
}

// CheckRules check if a pair of pages follows given rules
func (p Pair) CheckRules(rules Rules) bool {
	return !containsAny(rules[p[1]], p[0])
}

func containsAny[T comparable](input []T, elem T) bool {
	for _, v := range input {
		if v == elem {
			return true
		}
	}
	return false
}

func containsAll[T comparable](input []T, required []T) bool {
	for _, v := range required {
		if !containsAny(input, v) {
			return false
		}
	}
	return false
}

func orderedPair[T any](input []T) [][2]T {
	out := make([][2]T, 0)
	for i := 0; i+1 < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			out = append(out, [2]T{input[i], input[j]})
		}
	}
	return out
}
