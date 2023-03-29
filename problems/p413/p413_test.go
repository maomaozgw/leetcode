package p413

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
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
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "UE 1",
			args: args{
				nums: []int{1, 2, 3, 4, 4, 5, 6, 7, 8},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
