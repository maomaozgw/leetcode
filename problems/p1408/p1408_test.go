package p1408

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_stringMatching(t *testing.T) {
	type args struct {
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
				words: []string{"mass", "as", "hero", "superhero"},
			},
			want: []string{
				"as", "hero",
			},
		},
		{
			name: "Example 2",
			args: args{
				words: []string{"leetcode", "et", "code"},
			},
			want: []string{
				"et", "code",
			},
		},
		{
			name: "Example 3",
			args: args{
				words: []string{"blue", "green", "bu"},
			},
			want: []string{},
		},
		{
			name: "WA 1",
			args: args{
				words: []string{"leetcoder", "leetcode", "od", "hamlet", "am"},
			},
			want: []string{
				"leetcode", "od", "am",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringMatching(tt.args.words); !cmp.Equal(
				got, tt.want, cmpopts.SortSlices(func(a, b string) bool { return a < b }), cmpopts.EquateEmpty()) {
				t.Errorf("stringMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
