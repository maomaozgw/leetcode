package p1886

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_findRotation(t *testing.T) {
	type args struct {
		mat    [][]int
		target [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				mat:    tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[1,0]`),
				target: tools.NewGridFromStr[int](strconv.Atoi, `[1,0],[0,1]`),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				mat:    tools.NewGridFromStr[int](strconv.Atoi, `[0,1],[1,1]`),
				target: tools.NewGridFromStr[int](strconv.Atoi, `[1,0],[0,1]`),
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				mat:    tools.NewGridFromStr[int](strconv.Atoi, `[0,0,0],[0,1,0],[1,1,1]`),
				target: tools.NewGridFromStr[int](strconv.Atoi, `[1,1,1],[0,1,0],[0,0,0]`),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotation(tt.args.mat, tt.args.target); got != tt.want {
				t.Errorf("findRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
