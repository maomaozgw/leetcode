package p609

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Example 1",
			args: args{
				paths: []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			},
			want: [][]string{{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}},
		},
		{
			name: "Example 2",
			args: args{
				paths: []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"},
			},
			want: [][]string{
				{"root/a/2.txt", "root/c/d/4.txt"}, {"root/a/1.txt", "root/c/3.txt"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.paths); !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(i, j []string) bool {
				if len(i) == len(j) {
					return strings.Compare(i[0], j[0]) < 0
				}
				return len(i) < len(j)
			})) {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
