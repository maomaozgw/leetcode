package p1011

import "testing"

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		days    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				days:    5,
			},
			want: 15,
		},
		{
			name: "Example 1",
			args: args{
				weights: []int{
					3, 2, 2, 4, 1, 4,
				},
				days: 3,
			},
			want: 6,
		},
		{
			name: "Example 3",
			args: args{
				weights: []int{1, 2, 3, 1, 1},
				days:    4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.days); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
