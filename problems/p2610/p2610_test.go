package p2610

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_findMatrix(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 3, 4, 1, 2, 3, 1},
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[1,3,4,2],[1,3],[1]`),
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[1,2,3,4]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMatrix(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
