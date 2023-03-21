package p2348

import "testing"

func Test_zeroFilledSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 3, 0, 0, 2, 0, 0, 4},
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 0, 0, 2, 0, 0},
			},
			want: 9,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{2, 10, 2019},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroFilledSubarray(tt.args.nums); got != tt.want {
				t.Errorf("zeroFilledSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
