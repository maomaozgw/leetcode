package p587

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"reflect"
	"strconv"
	"testing"
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
			if got := outerTrees(tt.args.trees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outerTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
