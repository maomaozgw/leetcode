package p77

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				n: 4,
				k: 2,
			},
			want: tools.NewGridFromStr(strconv.Atoi, `[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]`),
		},
		{
			name: "Example 2",
			args: args{
				n: 1,
				k: 1,
			},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
