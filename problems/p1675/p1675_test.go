package p1675

import "testing"

func Test_minimumDeviation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{4, 1, 5, 20, 3},
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{2, 10, 8},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeviation(tt.args.nums); got != tt.want {
				t.Errorf("minimumDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
