package p823

import "testing"

func Test_numFactoredBinaryTrees(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				arr: []int{2, 4},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				arr: []int{2, 4, 5, 10},
			},
			want: 7,
		},
		{
			name: "WA ",
			args: args{
				arr: []int{18, 3, 6, 2},
			},
			want: 12,
		},
		{
			name: "WA 2",
			args: args{
				arr: []int{45, 42, 2, 18, 23, 1170, 12, 41, 40, 9, 47, 24, 33, 28, 10, 32, 29, 17, 46, 11, 759, 37, 6, 26, 21, 49, 31, 14, 19, 8, 13, 7, 27, 22, 3, 36, 34, 38, 39, 30, 43, 15, 4, 16, 35, 25, 20, 44, 5, 48},
			},
			want: 777,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numFactoredBinaryTrees(tt.args.arr); got != tt.want {
				t.Errorf("numFactoredBinaryTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
