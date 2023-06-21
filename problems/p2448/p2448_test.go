package p2448

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		nums []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 3, 5, 2},
				cost: []int{2, 3, 1, 14},
			},
			want: 8,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
				cost: []int{4, 2, 8, 1, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.nums, tt.args.cost); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
