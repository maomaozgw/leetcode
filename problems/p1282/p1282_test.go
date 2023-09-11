package p1282

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_groupThePeople(t *testing.T) {
	type args struct {
		groupSizes []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				groupSizes: []int{3, 3, 3, 3, 3, 1, 3},
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[0,1,2],[5],[3,4,6]`),
		},
		{
			name: "Example 2",
			args: args{
				groupSizes: []int{2, 1, 3, 3, 3, 2},
			},
			want: tools.NewGridFromStr[int](strconv.Atoi, `[[1],[2,3,4]],[0,5]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupThePeople(tt.args.groupSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupThePeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
