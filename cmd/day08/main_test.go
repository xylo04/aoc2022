package main

import "testing"

var exampleGrid = [][]uint8{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func Test_findNumVisibleTrees(t *testing.T) {
	type args struct {
		grid [][]uint8
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{grid: exampleGrid}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumVisibleTrees(tt.args.grid); got != tt.want {
				t.Errorf("findNumVisibleTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHidden(t *testing.T) {
	type args struct {
		grid [][]uint8
		r    int
		c    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "0,0", args: args{grid: exampleGrid, r: 0, c: 0}, want: false},
		{name: "1,1", args: args{grid: exampleGrid, r: 1, c: 1}, want: false},
		{name: "1,3", args: args{grid: exampleGrid, r: 1, c: 3}, want: true},
		{name: "2,1", args: args{grid: exampleGrid, r: 2, c: 1}, want: false},
		{name: "2,2", args: args{grid: exampleGrid, r: 2, c: 2}, want: true},
		{name: "4,4", args: args{grid: exampleGrid, r: 4, c: 4}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHidden(tt.args.grid, tt.args.r, tt.args.c); got != tt.want {
				t.Errorf("isHidden() = %v, want %v", got, tt.want)
			}
		})
	}
}
