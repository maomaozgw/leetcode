package p212

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/maomaozgw/leetcode/structures/tools"
	"strings"
	"testing"
)

func s2b(s string) (byte, error) {
	return strings.Trim(s, `"'`)[0], nil
}

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				board: tools.NewGridFromStr(s2b, `["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]`),
				words: []string{"oath", "pea", "eat", "rain"},
			},
			want: []string{"eat", "oath"},
		},
		{
			name: "Example 2",
			args: args{
				board: tools.NewGridFromStr(s2b, `["a","b"],["c","d"]`),
				words: []string{"abcb"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(s1, s2 string) bool {
				return strings.Compare(s1, s2) > 0
			})) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
