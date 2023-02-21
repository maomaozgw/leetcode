package p540

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
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
				nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{3, 3, 7, 7, 10, 11, 11},
			},
			want: 10,
		},
		{
			name: "WA 1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "WA 2",
			args: args{
				nums: []int{0, 1, 1, 2, 2, 5, 5},
			},
			want: 0,
		},
		{
			name: "UE 1",
			args: args{
				nums: []int{1, 1, 2, 2, 5, 5, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
