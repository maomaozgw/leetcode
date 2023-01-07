package p997

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"strconv"
	"testing"
)

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:     2,
				trust: tools.NewGridFromStr(strconv.Atoi, `[1,2]`),
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				n:     3,
				trust: tools.NewGridFromStr(strconv.Atoi, `[1,3],[2,3]`),
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				n:     3,
				trust: tools.NewGridFromStr(strconv.Atoi, `[1,3],[2,3],[3,1]`),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
