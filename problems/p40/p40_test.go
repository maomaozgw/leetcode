package p40

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[1,1,6],[1,2,5],[1,7],[2,6]`),
		},
		{
			name: "Example 2",
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[1,2,2],
			[5]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !tools.GridEqual(got, tt.want) {

				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
