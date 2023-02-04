package p1822

import "testing"

func Test_arraySign(t *testing.T) {
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
				nums: []int{-1, -2, -3, -4, 3, 2, 1},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{
					1, 5, 0, 2, -3,
				},
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{
					-1, 1, -1, 1, -1,
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arraySign(tt.args.nums); got != tt.want {
				t.Errorf("arraySign() = %v, want %v", got, tt.want)
			}
		})
	}
}
