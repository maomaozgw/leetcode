package p74

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
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
				matrix: tools.NewGridFromStr(strconv.Atoi, "[1,3,5,7],[10,11,16,20],[23,30,34,60]"),
				target: 3,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				matrix: tools.NewGridFromStr(strconv.Atoi, "[1,3,5,7],[10,11,16,20],[23,30,34,60]"),
				target: 13,
			},
			want: false,
		},
		{
			name: "UE 1",
			args: args{
				matrix: tools.NewGridFromStr(strconv.Atoi, "[1,3,5,7],[10,11,16,20],[23,30,34,60]"),
				target: 60,
			},
			want: true,
		},
		{
			name: "UE 2",
			args: args{
				matrix: tools.NewGridFromStr(strconv.Atoi, "[1,3,5,7],[10,11,16,20],[23,30,34,60]"),
				target: 1,
			},
			want: true,
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
