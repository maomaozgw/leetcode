package p1814

import (
	"testing"
)

func Test_countNicePairs(t *testing.T) {
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
				nums: []int{42, 11, 1, 97},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{13, 10, 35, 24, 76},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNicePairs(tt.args.nums); got != tt.want {
				t.Errorf("countNicePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rev(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				i: 120,
			},
			want: 21,
		},
		{
			name: "Example 2",
			args: args{
				i: 123,
			},
			want: 321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rev(tt.args.i); got != tt.want {
				t.Errorf("rev() = %v, want %v", got, tt.want)
			}
		})
	}
}
