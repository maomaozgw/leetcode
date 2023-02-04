package p33

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{
					4, 5, 6, 7, 0, 1, 2,
				},
				target: 3,
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findK(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums:  []int{1, 2, 3, 4, 5, 6},
				left:  0,
				right: 5,
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				nums:  []int{2, 3, 4, 5, 6, 1},
				left:  0,
				right: 5,
			},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{
					6, 1, 2, 3, 4, 5,
				},
				left:  0,
				right: 5,
			},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{
					3, 4, 5, 6, 1, 2,
				},
				left:  0,
				right: 5,
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{
					5, 6, 1, 2, 3, 4,
				},
				left:  0,
				right: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findK(tt.args.nums, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("findK() = %v, want %v", got, tt.want)
			}
		})
	}
}
