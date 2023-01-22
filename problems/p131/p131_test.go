package p131

import (
	"reflect"
	"strings"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Example 1",
			args: args{
				s: "aab",
			},
			want: tools.NewGridFromStr(func(s string) (string, error) { return strings.Trim(s, "\""), nil }, `["a","a","b"],["aa","b"]`),
		},
		{
			name: "Example 2",
			args: args{
				s: "a",
			},
			want: [][]string{
				{"a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
