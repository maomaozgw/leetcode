package p130

import (
	"reflect"
	"testing"

	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Example 1",
			args: args{
				board: tools.NewGridFromStr[byte](tools.ByteConverter, `["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]`),
			},
			want: tools.NewGridFromStr[byte](tools.ByteConverter, `["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]`),
		},
		{
			name: "Example 2",
			args: args{
				board: [][]byte{{'X'}},
			},
			want: [][]byte{{'X'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			if !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}
