package p427

import (
	"reflect"
	"testing"
)

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
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
				words: []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"},
			},
			want: []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"},
		},
		{
			name: "Example 2",
			args: args{
				words: []string{
					"cat", "dog", "catdog",
				},
			},
			want: []string{"catdog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllConcatenatedWordsInADict(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllConcatenatedWordsInADict() = %v, want %v", got, tt.want)
			}
		})
	}
}
