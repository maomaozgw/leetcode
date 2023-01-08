package p149

import (
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				points: tools.NewGridFromStr(strconv.Atoi, `[1,1],[2,2],[3,3]`),
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				points: tools.NewGridFromStr(strconv.Atoi, `[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]`),
			},
			want: 4,
		},
		{
			name: "WA 1",
			args: args{
				points: tools.NewGridFromStr(strconv.Atoi, `[0,1],[0,0],[0,4],[0,-2],[0,-1],[0,3],[0,-4]`),
			},
			want: 7,
		},
		{
			name: "WA 2",
			args: args{
				points: tools.NewGridFromStr(strconv.Atoi, `[0,0],[1,-1],[1,1]`),
			},
			want: 2,
		},
		{
			name: "WA 3",
			args: args{
				points: tools.NewGridFromStr(strconv.Atoi, `[-424,-512],[-4,-47],[0,-23],[-7,-65],[7,138],[0,27],[-5,-90],[-106,-146],[-420,-158],[-7,-128],[0,16],[-6,9],[-34,26],[-9,-166],[-570,-69],[-665,-85],[560,248],[1,-17],[630,277],[1,-7],[-287,-222],[30,250],[5,5],[-475,-53],[950,187],[7,-6],[-700,-274],[3,62],[-318,-390],[7,19],[-285,-21],[-5,4],[53,37],[-5,-1],[-2,-33],[-95,11],[4,1],[8,25],[700,306],[1,24],[-2,-6],[-35,-387],[-630,-245],[-328,-260],[-350,-129],[35,299],[-380,-37],[-9,-9],[210,103],[7,-5],[-3,-52],[-51,23],[-8,-147],[-371,-451],[-1,-14],[-41,6],[-246,-184],[350,161],[-212,-268],[-140,-42],[-9,-4],[-7,5],[10,6],[-15,-191],[-7,-4],[318,342],[-8,-71],[-68,20],[6,119],[6,13],[-280,-100],[140,74],[-760,-101],[0,-24],[-70,-13],[0,2],[0,-9],[106,98]`),
			},
			want: 14,
		},
		{
			name: "UT 1",
			args: args{
				points: tools.NewGridFromStr(strconv.Atoi, `[1,1],[2,2],[1,2],[2,3]`),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
