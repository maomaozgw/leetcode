package p1337

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_kWeakestRows(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				mat: tools.NewGridFromStr[int](strconv.Atoi, `[1,1,0,0,0],
			[1,1,1,1,0],
			[1,0,0,0,0],
			[1,1,0,0,0],
			[1,1,1,1,1]`),
				k: 3,
			},
			want: []int{2, 0, 3},
		},
		{
			name: "Example 2",
			args: args{
				mat: tools.NewGridFromStr[int](strconv.Atoi, `[1,0,0,0],
			[1,1,1,1],
			[1,0,0,0],
			[1,0,0,0]`),
				k: 2,
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kWeakestRows(tt.args.mat, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kWeakestRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
