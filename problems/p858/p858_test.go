package p858

import "testing"

func Test_mirrorReflection(t *testing.T) {
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				p: 2,
				q: 1,
			},
			want: 2,
		},
		{
			name: "Example 1",
			args: args{
				p: 3,
				q: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorReflection(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("mirrorReflection() = %v, want %v", got, tt.want)
			}
		})
	}
}
