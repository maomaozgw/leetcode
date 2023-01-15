package p2246

import "testing"

func Test_longestPath(t *testing.T) {
	type args struct {
		parent []int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				parent: []int{-1, 0, 0, 1, 1, 2},
				s:      "abacbe",
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				parent: []int{-1, 0, 0, 0},
				s:      "aabc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPath(tt.args.parent, tt.args.s); got != tt.want {
				t.Errorf("longestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
