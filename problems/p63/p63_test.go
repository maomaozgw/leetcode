package p63

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				obstacleGrid: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				obstacleGrid: [][]int{
					{0, 1},
					{0, 0},
				},
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				obstacleGrid: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{
				obstacleGrid: [][]int{
					{0, 1, 0, 0},
				},
			},
			want: 0,
		},
		{
			name: "WA 1",
			args: args{
				obstacleGrid: [][]int{
					{1},
				},
			},
			want: 0,
		},
		{
			name: "WA 2",
			args: args{
				obstacleGrid: [][]int{
					{0, 0},
					{0, 1},
				},
			},
			want: 0,
		},
		{
			name: "WA 3",
			args: args{
				obstacleGrid: [][]int{
					{1, 0},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
