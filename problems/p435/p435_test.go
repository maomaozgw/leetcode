package p435

import (
	"github.com/maomaozgw/leetcode/structures/tools"
	"strconv"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				intervals: tools.NewGridFromStr(strconv.Atoi, `[1,2],[2,3],[3,4],[1,3]`),
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				intervals: tools.NewGridFromStr(strconv.Atoi, `[1,2],[1,2],[1,2]`),
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				intervals: tools.NewGridFromStr(strconv.Atoi, `[1,2],[2,3]`),
			},
			want: 0,
		},
		{
			name: "WA 1",
			args: args{
				intervals: tools.NewGridFromStr(strconv.Atoi, `[-52,31],[-73,-26],[82,97],[-65,-11],[-62,-49],[95,99],[58,95],[-31,49],[66,98],[-63,2],[30,47],[-40,-26]`),
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
