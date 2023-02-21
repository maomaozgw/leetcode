package p1572

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_diagonalSum(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				mat: tools.NewGridFromStr[int](strconv.Atoi, `[1,2,3],[4,5,6],[7,8,9]`),
			},
			want: 25,
		},
		{
			name: "Example 2",
			args: args{
				mat: tools.NewGridFromStr[int](strconv.Atoi, `[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]`),
			},
			want: 8,
		},
		{
			name: "Example 3",
			args: args{
				mat: [][]int{{5}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalSum(tt.args.mat); got != tt.want {
				t.Errorf("diagonalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
