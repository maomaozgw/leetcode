package p378

import "testing"

func Test_kthSmallest(t *testing.T) {
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
					{1, 5, 9},
					{10, 11, 13},
					{12, 13, 15},
				},
				k: 8,
			},
			want: 13,
		},
		{
			name: "Example 1",
			args: args{
				matrix: [][]int{
					{1, 5, 9},
					{10, 11, 13},
					{12, 13, 15},
				},
				k: 3,
			},
			want: 9,
		},
		{
			name: "Example 1",
			args: args{
				matrix: [][]int{
					{1, 5, 9},
					{10, 11, 13},
					{12, 13, 15},
				},
				k: 6,
			},
			want: 12,
		},
		{
			name: "Example 2",
			args: args{
				matrix: [][]int{
					{-5},
				},
				k: 1,
			},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
