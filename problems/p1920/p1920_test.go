package p1920

import (
	"reflect"
	"strings"
	"testing"
)

func Test_buildArray(t *testing.T) {
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
				nums: []int{0, 2, 1, 5, 3, 4},
			},
			want: []int{0, 1, 2, 4, 5, 3},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{5, 0, 1, 2, 3, 4},
			},
			want: []int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		strings.Count()
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
