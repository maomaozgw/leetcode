package p871

import "testing"

func Test_minRefuelStops(t *testing.T) {
	type args struct {
		target    int
		startFuel int
		stations  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				target:    1,
				startFuel: 1,
				stations:  [][]int{},
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				target:    100,
				startFuel: 1,
				stations:  [][]int{{10, 100}},
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				target:    100,
				startFuel: 10,
				stations: [][]int{
					{10, 60},
					{20, 30},
					{30, 30},
					{60, 40},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRefuelStops(tt.args.target, tt.args.startFuel, tt.args.stations); got != tt.want {
				t.Errorf("minRefuelStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
