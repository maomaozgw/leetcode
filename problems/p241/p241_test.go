package p241

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				expression: "2-1-1",
			},
			want: []int{
				0, 2,
			},
		},
		{
			name: "Example 2",
			args: args{
				expression: "2*3-4*5",
			},
			want: []int{
				-34, -14, -10, -10, 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffWaysToCompute(tt.args.expression); !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(i, j int) bool {
				return i < j
			}), cmpopts.EquateEmpty(),
			) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
