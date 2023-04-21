package p633

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				c: 5,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				c: 3,
			},
			want: false,
		},
		{
			name: "UE 1",
			args: args{
				c: 4,
			},
			want: true,
		},
		{
			name: "UE 2",
			args: args{
				c: 8,
			},
			want: true,
		},
		{
			name: "WA 1",
			args: args{
				c: 1,
			},
			want: true,
		},
		{
			name: "TLE 2147483644",
			args: args{
				c: 2147483644,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
