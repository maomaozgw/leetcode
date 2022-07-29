package p1253gan

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func Test_reconstructMatrix(t *testing.T) {
	type args struct {
		upper  int
		lower  int
		colsum []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				upper:  2,
				lower:  1,
				colsum: []int{1, 1, 1},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				upper:  2,
				lower:  3,
				colsum: []int{2, 2, 1, 1},
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				upper:  5,
				lower:  5,
				colsum: []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1},
			},
			want: true,
		},
		{
			name: "WA 1",
			args: args{
				upper:  4,
				lower:  7,
				colsum: []int{2, 1, 2, 2, 1, 1, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructMatrix(tt.args.upper, tt.args.lower, tt.args.colsum); !validate(tt.args.upper, tt.args.lower, tt.args.colsum, got, tt.want) {
				t.Errorf("reconstructMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func validate(upper, lower int, colsum []int, result [][]int, hasResult bool) bool {
	if len(result) == 0 {
		return !hasResult
	}
	var (
		resultUpper, resultLower int
		resultColSum             = make([]int, len(result[0]))
	)
	for idx, val := range result[0] {
		resultUpper += val
		resultColSum[idx] += val
		resultLower += result[1][idx]
		resultColSum[idx] += result[1][idx]
	}
	if resultUpper != upper {
		return false
	}
	if resultLower != lower {
		return false
	}
	return cmp.Equal(resultColSum, colsum, cmpopts.EquateEmpty())
}
