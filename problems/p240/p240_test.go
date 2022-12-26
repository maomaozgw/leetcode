package p240

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"strconv"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				matrix: tools.NewGridFromStr(strconv.Atoi, `[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]`),
				target: 5,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				matrix: tools.NewGridFromStr(strconv.Atoi, `[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]`),
				target: 20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
