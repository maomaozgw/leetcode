package p763

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				s: "ababcbacadefegdehijhklij",
			},
			want: []int{9, 7, 8},
		},
		{
			name: "Example 2",
			args: args{
				s: "eccbbbbdec",
			},
			want: []int{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
