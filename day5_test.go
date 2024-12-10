package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseDay5(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name  string
		args  args
		want  Rules
		want1 []Document
	}{
		{name: "Minimal", args: args{lines: []string{"1|2", "", "1,2"}}, want: Rules{1: []int{2}}, want1: []Document{{Pages: []int{1, 2}}}},
		{name: "Multiple Docs", args: args{lines: []string{"1|2", "2|3", "", "1,2,3", "1,2", "1,3"}}, want: Rules{1: []int{2}, 2: []int{3}}, want1: []Document{{Pages: []int{1, 2, 3}}, {Pages: []int{1, 2}}, {Pages: []int{1, 3}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseDay5(tt.args.lines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDay5() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseDay5() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDocument_CheckRules(t *testing.T) {
	type args struct {
		rules Rules
	}
	tests := []struct {
		name string
		d    *Document
		args args
		want bool
	}{
		{name: "OK", d: &Document{Pages: []int{1, 2, 3}}, args: args{rules: Rules{1: {2, 3}}}, want: true},
		{name: "No Rules", d: &Document{Pages: []int{1, 2, 3}}, args: args{rules: Rules{1: {2, 3}}}, want: true},
		{name: "NOK", d: &Document{Pages: []int{2, 1, 3}}, args: args{rules: Rules{1: {2, 3}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.CheckRules(tt.args.rules)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestPair_CheckRules(t *testing.T) {
	tests := []struct {
		name  string
		p     Pair
		rules Rules
		want  bool
	}{
		{name: "OK", p: Pair{0, 1}, rules: Rules{0: []int{1, 2, 3}, 1: []int{2, 3}}, want: true},
		{name: "NOK", p: Pair{1, 0}, rules: Rules{0: []int{1, 2, 3}, 1: []int{2, 3}}, want: false},
		{name: "No Rules", p: Pair{0, 1}, rules: Rules{}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.CheckRules(tt.rules)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDocument_MovePage(t *testing.T) {
	tests := []struct {
		name string
		d    *Document
		pair Pair
		want []int
	}{
		{name: "OK", d: &Document{Pages: []int{1, 2, 3, 4, 5}}, pair: Pair{2, 5}, want: []int{1, 5, 2, 3, 4}},
		{name: "Edges", d: &Document{Pages: []int{1, 2, 3, 4, 5}}, pair: Pair{1, 5}, want: []int{5, 1, 2, 3, 4}},
		{name: "Small", d: &Document{Pages: []int{1, 5}}, pair: Pair{1, 5}, want: []int{5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.MovePage(tt.pair)
			require.Equal(t, tt.want, tt.d.Pages)
		})
	}
}

func TestDocument_Sort(t *testing.T) {
	tests := []struct {
		name  string
		d     *Document
		rules Rules
		want  []int
	}{
		{name: "OK", d: &Document{Pages: []int{1, 2, 3, 4, 5}}, rules: Rules{5: []int{2, 1, 3, 4}, 2: []int{1, 3, 4}, 1: []int{3, 4}, 3: []int{4}}, want: []int{5, 2, 1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Sort(tt.rules)
			require.Equal(t, tt.want, tt.d.Pages)
		})
	}
}
