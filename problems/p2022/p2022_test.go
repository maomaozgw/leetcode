package p2022

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_construct2DArray(t *testing.T) {
	type args struct {
		original []int
		m        int
		n        int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				original: []int{1, 2, 3, 4},
				m:        2,
				n:        2,
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[1,2],[3,4]`),
		},
		{
			name: "Example 2",
			args: args{
				original: []int{1, 2, 3},
				m:        1,
				n:        3,
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[1,2,3]`),
		},
		{
			name: "Example 3",
			args: args{
				original: []int{1, 2},
				m:        1,
				n:        1,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct2DArray(tt.args.original, tt.args.m, tt.args.n); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty()) {
				t.Errorf("construct2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
