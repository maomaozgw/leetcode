package p442

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			want: []int{2, 3},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: []int{1},
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty(), cmpopts.SortSlices(func(a, b int) bool { return a < b })) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
