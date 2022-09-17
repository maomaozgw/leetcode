package p336

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func Test_palindromePairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				words: []string{"abcd", "dcba", "lls", "s", "sssll"},
			},
			want: [][]int{
				{0, 1}, {1, 0}, {3, 2}, {2, 4},
			},
		},
		{
			name: "Example 2",
			args: args{
				words: []string{
					"bat", "tab", "cat",
				},
			},
			want: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name: "Example 3",
			args: args{
				words: []string{
					"a", "",
				},
			},
			want: [][]int{
				{0, 1}, {1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePairs(tt.args.words); !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(i, j int) bool { return i < j })) {
				t.Errorf("palindromePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
