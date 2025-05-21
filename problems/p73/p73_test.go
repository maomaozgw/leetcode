package p73

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
		want   [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Example 1",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
				want: [][]int{
					{1, 0, 1},
					{0, 0, 0},
					{1, 0, 1},
				},
			},
		},
		{
			name: "Example 2",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
				want: [][]int{
					{0, 0, 0, 0},
					{0, 4, 5, 0},
					{0, 3, 1, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.args.want) {
				t.Errorf("setZeroes() = %v, want %v", tt.args.matrix, tt.args.want)
			}
		})
	}
}
