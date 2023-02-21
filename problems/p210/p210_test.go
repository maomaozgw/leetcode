package p210

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				numCourses:    2,
				prerequisites: tools.NewGridFromStr[int](strconv.Atoi, `[1,0]`),
			},
			want: []int{0, 1},
		},
		{
			name: "Example 2",
			args: args{
				numCourses:    4,
				prerequisites: tools.NewGridFromStr[int](strconv.Atoi, `[1,0],[2,0],[3,1],[3,2]`),
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "Example 3",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
