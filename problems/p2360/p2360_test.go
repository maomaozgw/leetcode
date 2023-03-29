package p2360

import "testing"

func Test_longestCycle(t *testing.T) {
	type args struct {
		edges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				edges: []int{3, 3, 4, 2, 3},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				edges: []int{2, -1, 3, 1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCycle(tt.args.edges); got != tt.want {
				t.Errorf("longestCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
