package p363

import "testing"

func Test_maxSumSubmatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				matrix: [][]int{
					{1, 0, 1},
					{0, -2, 3},
				},
				k: 2,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				matrix: [][]int{
					{2, 2, -1},
				},
				k: 3,
			},
			want: 3,
		},
		{
			name: "WA 1",
			args: args{
				matrix: [][]int{
					{2, 2, -1},
				},
				k: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumSubmatrix(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("maxSumSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
