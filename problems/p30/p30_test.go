package p30

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				s:     "barfoothefoobarman",
				words: []string{"foo", "bar"},
			},
			want: []int{
				0, 9,
			},
		},
		{
			name: "Example 2",
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "word"},
			},
			want: []int{},
		},
		{
			name: "Example 3",
			args: args{
				s:     "barfoofoobarthefoobarman",
				words: []string{"bar", "foo", "the"},
			},
			want: []int{6, 9, 12},
		},
		{
			name: "Example 4",
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "good"},
			},
			want: []int{8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty()) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
