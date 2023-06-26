package p2462

import "testing"

func Test_totalCost(t *testing.T) {
	type args struct {
		costs      []int
		k          int
		candidates int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				costs:      []int{17, 12, 10, 2, 7, 2, 11, 20, 8},
				k:          3,
				candidates: 4,
			},
			want: 11,
		},
		{
			name: "Example 2",
			args: args{
				costs:      []int{1, 2, 4, 1},
				k:          3,
				candidates: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalCost(tt.args.costs, tt.args.k, tt.args.candidates); got != tt.want {
				t.Errorf("totalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
