package p1588

import "testing"

func Test_sumOddLengthSubarrays(t *testing.T) {
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
				arr: []int{
					1, 4, 2, 5, 3,
				},
			},
			want: 58,
		},
		{
			name: "Example 2",
			args: args{
				arr: []int{
					1, 2,
				},
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				arr: []int{
					10, 11, 12,
				},
			},
			want: 66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOddLengthSubarrays(tt.args.arr); got != tt.want {
				t.Errorf("sumOddLengthSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
