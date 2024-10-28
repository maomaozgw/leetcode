package p2501

import "testing"

func Test_longestSquareStreak(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{4, 3, 6, 16, 8, 2},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 3, 5, 6, 7},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSquareStreak(tt.args.nums); got != tt.want {
				t.Errorf("longestSquareStreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
