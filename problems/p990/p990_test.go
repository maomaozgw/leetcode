package p990

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				equations: []string{
					"a==b", "b!=a",
				},
			},
			want: false,
		},
		{
			name: "Example 2",
			args: args{
				equations: []string{
					"a==b", "b==a",
				},
			},
			want: true,
		},
		{
			name: "WA 1",
			args: args{
				equations: []string{
					"a==b", "e==c", "b==c", "a!=e",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
