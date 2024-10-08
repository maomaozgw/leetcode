package p1497

import "testing"

func Test_canArrange(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9},
				k:   5,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				arr: []int{
					1, 2, 3, 4, 5, 6,
				},
				k: 7,
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				arr: []int{
					1, 2, 3, 4, 5, 6,
				},
				k: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canArrange(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("canArrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
