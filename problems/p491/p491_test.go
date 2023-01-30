package p491

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_findSubsequences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{4, 6, 7, 7},
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]`),
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{4, 4, 3, 2, 1},
			},
			want: [][]int{
				{4, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubsequences(tt.args.nums); !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(i, j []int) bool {
				if len(i) == len(j) {
					for k := 0; k < len(i); k++ {
						if i[k] == j[k] {
							continue
						}
						return i[k] < j[k]
					}
					return true
				}
				return len(i) < len(j)
			})) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
