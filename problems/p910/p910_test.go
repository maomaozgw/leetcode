package p910

import "testing"

func Test_smallestRangeII(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1},
				k:    0,
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 10},
				k:    2,
			},
			want: 6,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 3, 6},
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRangeII(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestRangeII() = %v, want %v", got, tt.want)
			}
		})
	}
}
