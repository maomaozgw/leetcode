package p1861

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/maomaozgw/leetcode/structures/tools"
)

func Test_rotateTheBox(t *testing.T) {
	type args struct {
		box [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Example 1",
			args: args{
				box: tools.NewGridFromStr[byte](tools.ByteConverter, `["#",".","#"]`),
			},
			want: tools.NewGridFromStr[byte](tools.ByteConverter, `["."],
         ["#"],
         ["#"]`),
		},
		{
			name: "Example 2",
			args: args{
				box: tools.NewGridFromStr[byte](tools.ByteConverter, `["#",".","*","."],
              ["#","#","*","."]`),
			},
			want: tools.NewGridFromStr[byte](tools.ByteConverter, `["#","."],
         ["#","#"],
         ["*","*"],
         [".","."]`),
		},
		{
			name: "Example 3",
			args: args{
				box: tools.NewGridFromStr[byte](tools.ByteConverter, `["#","#","*",".","*","."],
              ["#","#","#","*",".","."],
              ["#","#","#",".","#","."]`),
			},
			want: tools.NewGridFromStr[byte](tools.ByteConverter, `[".","#","#"],
         [".","#","#"],
         ["#","#","*"],
         ["#","*","."],
         ["#",".","*"],
         ["#",".","."]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateTheBox(tt.args.box); !cmp.Equal(got, tt.want) {
				t.Errorf("rotateTheBox() = %v, want %v, diff: %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
