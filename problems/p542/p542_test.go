package p542

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				mat: tools.NewGridFromStr(strconv.Atoi, `[0,0,0],[0,1,0],[0,0,0]`),
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[0,0,0],[0,1,0],[0,0,0]`),
		},
		{
			name: "Example 2",
			args: args{
				mat: tools.NewGridFromStr(strconv.Atoi, `[0,0,0],[0,1,0],[1,1,1]`),
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[0,0,0],[0,1,0],[1,2,1]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
