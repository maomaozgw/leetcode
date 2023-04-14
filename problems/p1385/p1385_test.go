package p1385

import "testing"

func Test_findTheDistanceValue(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		d    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				arr1: []int{4, 5, 8},
				arr2: []int{10, 9, 1, 8},
				d:    2,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				arr1: []int{1, 4, 2, 3},
				arr2: []int{-4, -3, 6, 10, 20, 30},
				d:    3,
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				arr1: []int{2, 1, 100, 3},
				arr2: []int{-5, -2, 10, -3, 7},
				d:    6,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDistanceValue(tt.args.arr1, tt.args.arr2, tt.args.d); got != tt.want {
				t.Errorf("findTheDistanceValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
