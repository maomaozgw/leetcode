package p931

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"strconv"
	"testing"
)

func Test_minFallingPathSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				matrix: tools.NewGridFromStr(strconv.Atoi, "[2,1,3],[6,5,4],[7,8,9]"),
			},
			want: 13,
		},
		{
			name: "Example 2",
			args: args{
				matrix: tools.NewGridFromStr(strconv.Atoi, "[-19,57],[-40,-5]"),
			},
			want: -59,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
