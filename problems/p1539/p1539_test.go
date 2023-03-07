package p1539

import "testing"

func Test_findKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				arr: []int{2, 3, 4, 7, 11},
				k:   5,
			},
			want: 9,
		},
		{
			name: "Example 2",
			args: args{
				arr: []int{1, 2, 3, 4},
				k:   2,
			},
			want: 6,
		},
		{
			name: "Example 3",
			args: args{
				arr: []int{7, 8, 9, 10},
				k:   6,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthPositive(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findKthPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
