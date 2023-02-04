package p1790

import "testing"

func Test_areAlmostEqual(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s1: "bank",
				s2: "kanb",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s1: "attack",
				s2: "defend",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				s1: "kelb",
				s2: "kelb",
			},
			want: true,
		},
		{
			name: "WA 1",
			args: args{
				s1: "caa",
				s2: "aaz",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areAlmostEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("areAlmostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
