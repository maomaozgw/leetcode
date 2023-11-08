package p2849

import "testing"

func Test_isReachableAtTime(t *testing.T) {
	type args struct {
		sx int
		sy int
		fx int
		fy int
		t  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				sx: 2,
				sy: 4,
				fx: 7,
				fy: 7,
				t:  6,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				sx: 3,
				sy: 1,
				fx: 7,
				fy: 3,
				t:  3,
			},
			want: false,
		},
		{
			name: "Same Point",
			args: args{
				sx: 3,
				sy: 1,
				fx: 3,
				fy: 1,
				t:  1,
			},
			want: false,
		},
		{
			name: "WA 1",
			args: args{
				sx: 1,
				sy: 1,
				fx: 4,
				fy: 1,
				t:  0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReachableAtTime(tt.args.sx, tt.args.sy, tt.args.fx, tt.args.fy, tt.args.t); got != tt.want {
				t.Errorf("isReachableAtTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
