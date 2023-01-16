package p57

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				intervals:   tools.NewGridFromStr(strconv.Atoi, `[1,3],[6,9]`),
				newInterval: []int{2, 5},
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[1,5],[6,9]`),
		},
		{
			name: "Example 2",
			args: args{
				intervals:   tools.NewGridFromStr(strconv.Atoi, `[1,2],[3,5],[6,7],[8,10],[12,16]`),
				newInterval: []int{4, 8},
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[1,2],[3,10],[12,16]`),
		},
		{
			name: "UE 1",
			args: args{
				intervals:   tools.NewGridFromStr(strconv.Atoi, `[1,2],[3,5],[6,7],[8,10],[12,16]`),
				newInterval: []int{3, 4},
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[1,2],[3,5],[6,7],[8,10],[12,16]`),
		},
		{
			name: "UE 2",
			args: args{
				intervals:   tools.NewGridFromStr(strconv.Atoi, `[1,2],[3,5],[6,7],[8,10],[12,16]`),
				newInterval: []int{15, 18},
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[1,2],[3,5],[6,7],[8,10],[12,18]`),
		},
		{
			name: "UE 3",
			args: args{
				intervals:   tools.NewGridFromStr(strconv.Atoi, `[1,2],[3,5],[6,7],[8,10],[12,16]`),
				newInterval: []int{17, 18},
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[1,2],[3,5],[6,7],[8,10],[12,16],[17,18]`),
		},
		{
			name: "UE 3",
			args: args{
				intervals:   tools.NewGridFromStr(strconv.Atoi, `[3,5],[6,7],[8,10],[12,16]`),
				newInterval: []int{1, 2},
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[1,2],[3,5],[6,7],[8,10],[12,16]`),
		},
		{
			name: "Wa 1",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{0, 3},
			},
			want: [][]int{{0, 5}},
		},
		{
			name: "UE 4",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{0, 6},
			},
			want: [][]int{{0, 6}},
		},
		{
			name: "WA 2",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{0, 1},
			},
			want: [][]int{{0, 5}},
		},
		{
			name: "UE 5",
			args: args{
				intervals:   nil,
				newInterval: []int{0, 5},
			},
			want: [][]int{{0, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
