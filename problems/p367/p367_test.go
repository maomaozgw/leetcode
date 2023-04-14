package p367

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				num: 16,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				num: 14,
			},
			want: false,
		},
		{
			name: "UE 1",
			args: args{
				num: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
