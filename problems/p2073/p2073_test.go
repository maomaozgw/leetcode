package p2073

import "testing"

func Test_timeRequiredToBuy(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				tickets: []int{2, 3, 2},
				k:       2,
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				tickets: []int{5, 1, 1, 1},
				k:       0,
			},
			want: 8,
		},
		{
			name: "WA 1",
			args: args{
				tickets: []int{84, 49, 5, 24, 70, 77, 87, 8},
				k:       3,
			},
			want: 154,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeRequiredToBuy(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("timeRequiredToBuy() = %v, want %v", got, tt.want)
			}
		})
	}
}
