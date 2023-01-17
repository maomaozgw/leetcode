package p936

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_movesToStamp(t *testing.T) {
	type args struct {
		stamp  string
		target string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				stamp:  "abc",
				target: "ababc",
			},
			want: []int{0, 2},
		},
		{
			name: "Example 2",
			args: args{
				stamp:  "abca",
				target: "aabcaca",
			},
			want: []int{3, 0, 1},
		},
		{
			name: "WA 1",
			args: args{
				stamp:  "de",
				target: "ddeddeddee",
			},
			want: []int{6, 3, 0, 8, 7, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movesToStamp(tt.args.stamp, tt.args.target); !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(i, j int) bool {
				return i < j
			})) {
				t.Errorf("movesToStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
