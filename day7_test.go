package main

import (
	"reflect"
	"testing"
)

func Test_tryOp(t *testing.T) {
	type args struct {
		curValue int
		values   []int
		ops      []Op
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Simple", args: args{curValue: 2, values: []int{2}, ops: []Op{opAdd}}, want: []int{4}},
		{name: "Duplicate", args: args{curValue: 2, values: []int{2}, ops: []Op{opAdd, opMult}}, want: []int{4, 4}},
		{name: "Empty Ops", args: args{curValue: 2, values: []int{2}, ops: []Op{}}, want: []int{}},
		{name: "Empty Values", args: args{curValue: 2, values: []int{}, ops: []Op{opAdd, opMult}}, want: []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tryOp(tt.args.curValue, tt.args.values, tt.args.ops); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tryOp() = %v, want %v", got, tt.want)
			}
		})
	}
}
