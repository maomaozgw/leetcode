package p260

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_singleNumber(t *testing.T) {
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
				nums: []int{1, 2, 1, 3, 2, 5},
			},
			want: []int{3, 5},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{-1, 0},
			},
			want: []int{-1, 0},
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{0, 1},
			},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(s1, s2 int) bool {
				return s1 > s2
			})) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
