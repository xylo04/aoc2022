package main

import "testing"

func Test_findMarker(t *testing.T) {
	type args struct {
		signal string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{signal: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, want: 7},
		{name: "1", args: args{signal: "bvwbjplbgvbhsrlpgdmjqwftvncz"}, want: 5},
		{name: "2", args: args{signal: "nppdvjthqldpwncqszvftbrmjlhg"}, want: 6},
		{name: "3", args: args{signal: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, want: 10},
		{name: "4", args: args{signal: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMarker(tt.args.signal); got != tt.want {
				t.Errorf("findMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
