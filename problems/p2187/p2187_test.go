package p2187

import "testing"

func Test_minimumTime(t *testing.T) {
	type args struct {
		time       []int
		totalTrips int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				time:       []int{1, 2, 3},
				totalTrips: 5,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				time:       []int{2},
				totalTrips: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime(tt.args.time, tt.args.totalTrips); got != tt.want {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
