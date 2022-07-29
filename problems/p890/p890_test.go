package p890

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func Test_findAndReplacePattern(t *testing.T) {
	type args struct {
		words   []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				words:   []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
				pattern: "abb",
			},
			want: []string{"mee", "aqq"},
		},
		{
			name: "Example 2",
			args: args{
				words:   []string{"a", "b", "c"},
				pattern: "a",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "WA 1",
			args: args{
				words:   []string{"abc", "cba", "xyx", "yxx", "yyx"},
				pattern: "abc",
			},
			want: []string{"abc", "cba"},
		},
		{
			name: "WA 2",
			args: args{
				words:   []string{"ktittgzawn", "dgphvfjniv", "gceqobzmis", "alrztxdlah", "jijuevoioe", "mawiizpkub", "onwpmnujos", "zszkptjgzj", "zwfvzhrucv", "isyaphcszn"},
				pattern: "zdqmjnczma",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAndReplacePattern(tt.args.words, tt.args.pattern); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty()) {
				t.Errorf("findAndReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
