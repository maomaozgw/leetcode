package p47

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_permuteUnique(t *testing.T) {
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
				nums: []int{1, 1, 2},
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[1,1,2],[1,2,1],[2,1,1]`),
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,2,1],[3,1,2]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(a, b []int) bool {
				for i := range a {
					if a[i] == b[i] {
						continue
					} else if a[i] < b[i] {
						return true
					}
					return false
				}
				return false
			})) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
