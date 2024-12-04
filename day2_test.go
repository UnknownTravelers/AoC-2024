package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_run2a(t *testing.T) {
	type args struct {
		reports []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "OK Increase", args: args{reports: []string{"0 1 2 3 4 5"}}, want: 1},
		{name: "OK Decrease", args: args{reports: []string{"5 4 3 2 1 0"}}, want: 1},
		{name: "Not OK Growth Direction", args: args{reports: []string{"5 4 3 5 1 0"}}, want: 0},
		{name: "Not OK Big Step", args: args{reports: []string{"10 5 4 3 2 1 0"}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := run2a(tt.args.reports)
			require.NoError(t, err)
			require.Equal(t, tt.want, actual)
		})
	}
}
