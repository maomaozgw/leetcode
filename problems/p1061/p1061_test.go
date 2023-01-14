package p1061

import "testing"

func Test_smallestEquivalentString(t *testing.T) {
	type args struct {
		s1      string
		s2      string
		baseStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s1:      "parker",
				s2:      "morris",
				baseStr: "parser",
			},
			want: "makkek",
		},
		{
			name: "Example 2",
			args: args{
				s1:      "hello",
				s2:      "world",
				baseStr: "hold",
			},
			want: "hdld",
		},
		{
			name: "Example 3",
			args: args{
				s1:      "leetcode",
				s2:      "programs",
				baseStr: "sourcecode",
			},
			want: "aauaaaaada",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestEquivalentString(tt.args.s1, tt.args.s2, tt.args.baseStr); got != tt.want {
				t.Errorf("smallestEquivalentString() = %v, want %v", got, tt.want)
			}
		})
	}
}
