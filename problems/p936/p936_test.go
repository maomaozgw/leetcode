package p936

import (
	"reflect"
	"testing"
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
			want: []int{5, 2, 8, 6, 3, 0, 7, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movesToStamp(tt.args.stamp, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movesToStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
