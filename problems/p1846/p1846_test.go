package p1846

import "testing"

func Test_maximumElementAfterDecrementingAndRearranging(t *testing.T) {
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
				arr: []int{2, 2, 1, 2, 1},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				arr: []int{100, 1, 1000},
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "WA 1",
			args: args{
				arr: []int{73, 98, 9},
			},
			want: 3,
		},
		{
			name: "UC 1",
			args: args{
				arr: []int{99},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumElementAfterDecrementingAndRearranging(tt.args.arr); got != tt.want {
				t.Errorf("maximumElementAfterDecrementingAndRearranging() = %v, want %v", got, tt.want)
			}
		})
	}
}
