package p409

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "abccccdd",
			},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "UC 1",
			args: args{
				s: "zzzzzAAAAA",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
