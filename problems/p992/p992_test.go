package p992

import "testing"

func Test_subarraysWithKDistinct(t *testing.T) {
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
			name: "Example 1 ",
			args: args{
				nums: []int{1, 2, 1, 2, 3},
				k:    2,
			},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 2, 1, 3, 4},
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
