package p2657

import (
	"reflect"
	"testing"
)

func Test_findThePrefixCommonArray(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				A: []int{1, 3, 2, 4},
				B: []int{3, 1, 2, 4},
			},
			want: []int{0, 2, 3, 4},
		},
		{
			name: "Example 2",
			args: args{
				A: []int{2, 3, 1},
				B: []int{3, 1, 2},
			},
			want: []int{0, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThePrefixCommonArray(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findThePrefixCommonArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
