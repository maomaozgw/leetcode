package p34

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "example 2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
		{
			name: "example 3",
			args: args{
				nums:   []int{},
				target: 8,
			},
			want: []int{-1, -1},
		},
		{
			name: "wa 1",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: []int{0, 0},
		},
		{
			name: "wa 2",
			args: args{
				nums:   []int{1, 4},
				target: 4,
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
