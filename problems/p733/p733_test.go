package p733

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image [][]int
		sr    int
		sc    int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				image: tools.NewGridFromStr(strconv.Atoi, `[1,1,1],[1,1,0],[1,0,1]`),
				sr:    1,
				sc:    1,
				color: 2,
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[2,2,2],[2,2,0],[2,0,1]`),
		},
		{
			name: "Example 2",
			args: args{
				image: tools.NewGridFromStr(strconv.Atoi, `[0,0,0],[0,0,0]`),
				sr:    0,
				sc:    0,
				color: 0,
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[0,0,0],[0,0,0]`),
		},
		{
			name: "WA 1",
			args: args{
				image: tools.NewGridFromStr(strconv.Atoi, `[[0,0,0],[0,0,0]]`),
				sr:    1,
				sc:    0,
				color: 2,
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[[2,2,2],[2,2,2]]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
