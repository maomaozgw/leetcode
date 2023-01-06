package p1833

import "testing"

func Test_maxIceCream(t *testing.T) {
	type args struct {
		costs []int
		coins int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				costs: []int{1, 3, 2, 4, 1},
				coins: 7,
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				costs: []int{10, 6, 8, 7, 7, 8},
				coins: 5,
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				costs: []int{1, 6, 3, 1, 2, 5},
				coins: 20,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIceCream(tt.args.costs, tt.args.coins); got != tt.want {
				t.Errorf("maxIceCream() = %v, want %v", got, tt.want)
			}
		})
	}
}
