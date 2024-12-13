package p2593

import "testing"

func Test_findScore(t *testing.T) {
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
				nums: []int{2, 1, 3, 4, 5, 2},
			},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 3, 5, 1, 3, 2},
			},
			want: 5,
		},
		{
			name: "WA 1",
			args: args{
				nums: []int{3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findScore(tt.args.nums); got != tt.want {
				t.Errorf("findScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
