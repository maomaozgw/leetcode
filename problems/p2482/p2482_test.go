package p2482

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_onesMinusZeros(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[0,1,1],[1,0,1],[0,0,1]`),
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[0,0,4],[0,0,4],[-2,-2,2]`),
		},
		{
			name: "Example 2",
			args: args{
				grid: tools.NewGridFromStr[int](strconv.Atoi, `[1,1,1],[1,1,1]`),
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[5,5,5],[5,5,5]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := onesMinusZeros(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("onesMinusZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
