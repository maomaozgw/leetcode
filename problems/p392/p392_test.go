package p392

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
		{
			name: "WA 1",
			args: args{
				s: "acb",
				t: "ahbdgc",
			},
			want: false,
		},
		{
			name: "WA 2",
			args: args{
				s: "aaaaaa",
				t: "bbaaaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
