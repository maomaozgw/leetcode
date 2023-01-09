package p973

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				points: tools.NewGridFromStr(strconv.Atoi, `[1,3],[-2,2]`),
				k:      1,
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[-2,2]`),
		},
		{
			name: "Example 2",
			args: args{
				points: tools.NewGridFromStr(strconv.Atoi, `[3,3],[5,-1],[-2,4]`),
				k:      2,
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[3,3],[-2,4]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
