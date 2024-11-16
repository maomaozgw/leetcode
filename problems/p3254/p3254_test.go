package p3254

import (
	"reflect"
	"testing"
)

func Test_resultsArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{
					1, 2, 3, 4, 3, 2, 5,
				},
				k: 3,
			},
			want: []int{3, 4, -1, -1, -1},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
				k:    4,
			},
			want: []int{-1, -1},
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{3, 2, 3, 2, 3, 2},
				k:    2,
			},
			want: []int{
				-1, 3, -1, 3, -1,
			},
		},
		{
			name: "WA 1",
			args: args{
				nums: []int{1, 3, 4},
				k:    2,
			},
			want: []int{-1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resultsArray(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resultsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
