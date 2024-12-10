package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseInts(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Empty", args: args{str: "", sep: ""}, want: []int{}},
		{name: "1 number", args: args{str: "1", sep: ","}, want: []int{1}},
		{name: "1 neg number", args: args{str: "-1", sep: ","}, want: []int{-1}},
		{name: "Nominal", args: args{str: "1 2 3 4 5", sep: " "}, want: []int{1, 2, 3, 4, 5}},
		{name: "Nominal Large", args: args{str: "-500 -400 -300 -200 -100 0 100 200 300 400 500", sep: " "}, want: []int{-500, -400, -300, -200, -100, 0, 100, 200, 300, 400, 500}},
		{name: "Large", args: args{str: "99999999999999999", sep: ","}, want: []int{99999999999999999}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInts(tt.args.str, tt.args.sep)
			require.Equal(t, tt.want, got)
		})
	}
}
