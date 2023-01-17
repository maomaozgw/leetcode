package p587

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_outerTrees(t *testing.T) {
	type args struct {
		trees [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				trees: tools.NewGridFromStr(strconv.Atoi, "[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]"),
			},
			want: tools.NewGridFromStr(strconv.Atoi, "[1,1],[2,0],[3,3],[2,4],[4,2]"),
		},
		{
			name: "Example 2",
			args: args{
				trees: tools.NewGridFromStr(strconv.Atoi, "[1,2],[2,2],[4,2]"),
			},
			want: tools.NewGridFromStr(strconv.Atoi, "[4,2],[2,2],[1,2]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outerTrees(tt.args.trees); !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(i, j []int) bool {
				if i[0] == j[0] {
					return i[1] < j[1]
				}
				return i[0] < j[0]
			})) {
				t.Errorf("outerTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
