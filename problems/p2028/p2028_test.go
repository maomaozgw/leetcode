package p2028

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_missingRolls(t *testing.T) {
	type args struct {
		rolls []int
		mean  int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				rolls: []int{3, 2, 4, 3},
				mean:  4,
				n:     2,
			},
			want: []int{6, 6},
		},
		{
			name: "Exmaple 2",
			args: args{
				rolls: []int{1, 5, 6},
				mean:  3,
				n:     4,
			},
			want: []int{2, 3, 2, 2},
		},
		{
			name: "Example 3",
			args: args{
				rolls: []int{1, 2, 3, 4},
				mean:  6,
				n:     4,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingRolls(tt.args.rolls, tt.args.mean, tt.args.n); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty(), cmpopts.SortSlices(func(a, b int) bool { return a < b })) {
				t.Errorf("missingRolls() = %v, want %v", got, tt.want)
			}
		})
	}
}
