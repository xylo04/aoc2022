package main

import (
	"reflect"
	"testing"
)

func Test_tailVisitCount(t *testing.T) {
	type args struct {
		vectors []vector
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "1",
			args: args{
				[]vector{
					{"R", 4},
					{"U", 4},
					{"L", 3},
					{"D", 1},
					{"R", 4},
					{"D", 1},
					{"L", 5},
					{"R", 2},
				},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tailVisitCount(tt.args.vectors); got != tt.want {
				t.Errorf("tailVisitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_step(t *testing.T) {
	type args struct {
		head      location
		tail      location
		direction string
	}
	tests := []struct {
		name  string
		args  args
		wHead location
		wTail location
	}{
		{
			name:  "U1",
			args:  args{location{0, 0}, location{0, 0}, "U"},
			wHead: location{0, 1},
			wTail: location{0, 0},
		},
		{
			name:  "U2",
			args:  args{location{0, 1}, location{0, 0}, "U"},
			wHead: location{0, 2},
			wTail: location{0, 1},
		},
		{
			name:  "U3",
			args:  args{location{1, 1}, location{0, 0}, "U"},
			wHead: location{1, 2},
			wTail: location{1, 1},
		},
		{
			name:  "U4",
			args:  args{location{1, 1}, location{0, 2}, "U"},
			wHead: location{1, 2},
			wTail: location{0, 2},
		},
		{
			name:  "R1",
			args:  args{location{0, 0}, location{0, 0}, "R"},
			wHead: location{1, 0},
			wTail: location{0, 0},
		},
		{
			name:  "R2",
			args:  args{location{1, 0}, location{0, 0}, "R"},
			wHead: location{2, 0},
			wTail: location{1, 0},
		},
		{
			name:  "R3",
			args:  args{location{1, 1}, location{0, 0}, "R"},
			wHead: location{2, 1},
			wTail: location{1, 1},
		},
		{
			name:  "R4",
			args:  args{location{1, 1}, location{0, 2}, "R"},
			wHead: location{2, 1},
			wTail: location{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gHead, gTail := step(tt.args.head, tt.args.tail, tt.args.direction)
			if !reflect.DeepEqual(gHead, tt.wHead) {
				t.Errorf("step() gHead = %v, wHead %v", gHead, tt.wHead)
			}
			if !reflect.DeepEqual(gTail, tt.wTail) {
				t.Errorf("step() gTail = %v, wTail %v", gTail, tt.wTail)
			}
		})
	}
}
