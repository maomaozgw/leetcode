package p1608

import "testing"

func Test_specialArray(t *testing.T) {
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
				nums: []int{3, 5},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 0},
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{0, 4, 3, 0, 4},
			},
			want: 3,
		},
		{
			name: "WA 1",
			args: args{
				nums: []int{2, 3, 0, 2},
			},
			want: -1,
		},
		{
			name: "WA 2",
			args: args{
				nums: []int{3, 9, 7, 8, 3, 8, 6, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialArray(tt.args.nums); got != tt.want {
				t.Errorf("specialArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
