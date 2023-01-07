package p1557

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"reflect"
	"strconv"
	"testing"
)

func Test_findSmallestSetOfVertices(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				n:     6,
				edges: tools.NewGridFromStr(strconv.Atoi, `[0,1],[0,2],[2,5],[3,4],[4,2]`),
			},
			want: []int{0, 3},
		},
		{
			name: "Example 2",
			args: args{
				n:     5,
				edges: tools.NewGridFromStr(strconv.Atoi, `[0,1],[2,1],[3,1],[1,4],[2,4]`),
			},
			want: []int{0, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSmallestSetOfVertices(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSmallestSetOfVertices() = %v, want %v", got, tt.want)
			}
		})
	}
}
