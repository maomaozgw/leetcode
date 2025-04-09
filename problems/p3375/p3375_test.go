package p3375

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{5, 2, 5, 4, 5},
				k:    2,
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{2, 1, 2},
				k:    2,
			},
			want: -1,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{9, 7, 5, 3},
				k:    1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
