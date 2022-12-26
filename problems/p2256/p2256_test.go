package p2256

import "testing"

func Test_minimumAverageDifference(t *testing.T) {
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
				nums: []int{2, 5, 3, 9, 5, 3},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAverageDifference(tt.args.nums); got != tt.want {
				t.Errorf("minimumAverageDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
