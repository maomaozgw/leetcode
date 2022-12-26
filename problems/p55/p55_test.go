package p55

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{3, 2, 1, 0, 4},
			},
			want: false,
		},
		{
			name: "WA 1",
			args: args{
				nums: []int{0},
			},
			want: true,
		},
		{
			name: "UD 1",
			args: args{
				nums: []int{1, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
