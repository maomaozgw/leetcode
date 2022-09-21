package p985

import (
	"reflect"
	"testing"
)

func Test_sumEvenAfterQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 2",
			args: args{
				nums:    []int{1},
				queries: [][]int{{4, 0}},
			},
			want: []int{0},
		},
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4},
				queries: [][]int{
					{1, 0},
					{-3, 1},
					{-4, 0},
					{2, 3},
				},
			},
			want: []int{8, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumEvenAfterQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumEvenAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
