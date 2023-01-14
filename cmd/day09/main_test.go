package main

import "testing"

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
