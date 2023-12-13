package p1582

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_numSpecial(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				mat: tools.NewGridFromStr[int](strconv.Atoi, `[1,0,0],[0,0,1],[1,0,0]`),
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				mat: tools.NewGridFromStr[int](strconv.Atoi, `[1,0,0],[0,1,0],[0,0,1]`),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSpecial(tt.args.mat); got != tt.want {
				t.Errorf("numSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
