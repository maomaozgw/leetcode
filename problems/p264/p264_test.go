package p264

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n: 10,
			},
			want: 12,
		},
		{
			name: "Example 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "WA 1",
			args: args{
				n: 11,
			},
			want: 15,
		},
		{
			name: "WA 2",
			args: args{
				n: 1690,
			},
			want: 2123366400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
