package p1706

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_findBall(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				grid: tools.NewGridFromStr(strconv.Atoi, "[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-1,-1]"),
			},
			want: []int{1, -1, -1, -1, -1},
		},
		{
			name: "Example 2",
			args: args{
				grid: tools.NewGridFromStr(strconv.Atoi, "[-1]"),
			},
			want: []int{-1},
		},
		{
			name: "Example 3",
			args: args{
				grid: tools.NewGridFromStr(strconv.Atoi, "[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1],[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1]"),
			},
			want: []int{0, 1, 2, 3, 4, -1},
		},
		{
			name: "WA 1",
			args: args{
				grid: tools.NewGridFromStr(strconv.Atoi, "[-1,1,-1,-1,-1,-1,-1,-1,1,-1,-1,-1,-1,1,1,-1,-1,-1,1,1,1,-1,-1,1,1,-1,-1,1,-1,-1,-1,-1,-1,-1,-1,-1,-1,1,-1,1,-1,-1,-1,-1,-1,-1,-1,1,-1,-1,1,-1,1,-1,-1,1,1,-1,1,-1,-1,-1,-1,1,1,1,1,1,1,-1,1,1,1,-1,1,1,1,-1,-1,-1,1,-1,1,-1,-1,1,1,-1,-1,1,-1,1,-1,1,1,1,-1,-1,-1,-1]"),
			},
			want: []int{-1, -1, -1, 2, 3, 4, 5, 6, -1, -1, 9, 10, 11, 14, -1, -1, 15, 16, 19, 20, -1, -1, 21, 24, -1, -1, 25, -1, -1, 28, 29, 30, 31, 32, 33, 34, 35, -1, -1, -1, -1, 40, 41, 42, 43, 44, 45, -1, -1, 48, -1, -1, -1, -1, 53, 56, -1, -1, -1, -1, 59, 60, 61, 64, 65, 66, 67, 68, -1, -1, 71, 72, -1, -1, 75, 76, -1, -1, 77, 78, -1, -1, -1, -1, 83, 86, -1, -1, 87, -1, -1, -1, -1, 94, 95, -1, -1, 96, 97, 98},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBall(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBall() = %v, want %v", got, tt.want)
			}
		})
	}
}
