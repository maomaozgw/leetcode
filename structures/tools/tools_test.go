package tools

import (
	"strconv"
	"testing"
)

func TestNewGridFromStr(t *testing.T) {
	type args struct {
		convert func(string) (int, error)
		s       string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "with empty slice",
			args: args{
				convert: strconv.Atoi,
				s:       "[],[1],[2],[1,2]",
			},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGridFromStr(tt.args.convert, tt.args.s); !GridEqual(got, tt.want) {
				t.Errorf("NewGridFromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
