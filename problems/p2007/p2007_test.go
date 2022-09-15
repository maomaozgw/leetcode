package p2007

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func Test_findOriginalArray(t *testing.T) {
	type args struct {
		changed []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				changed: []int{1, 3, 4, 2, 6, 8},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "Example 2",
			args: args{
				changed: []int{6, 3, 0, 1},
			},
			want: []int{},
		},
		{
			name: "Example 3",
			args: args{
				changed: []int{1},
			},
			want: []int{},
		},
		{
			name: "WA 1",
			args: args{
				changed: []int{0, 0, 0, 0},
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOriginalArray(tt.args.changed); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty(), cmpopts.SortSlices(func(i, j int) bool { return i < j })) {
				t.Errorf("findOriginalArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
