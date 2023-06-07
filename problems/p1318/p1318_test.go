package p1318

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				a: 2,
				b: 6,
				c: 5,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				a: 4,
				b: 2,
				c: 7,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				a: 1,
				b: 2,
				c: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
