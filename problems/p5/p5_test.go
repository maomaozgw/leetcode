package p5

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "Example 2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "UE 1",
			args: args{
				s: "abcd",
			},
			want: "a",
		},
		{
			name: "UE 2",
			args: args{
				s: "abdd",
			},
			want: "dd",
		},
		{
			name: "WA",
			args: args{
				s: "bb",
			},
			want: "bb",
		},
		{
			name: "UE 3",
			args: args{
				s: "",
			},
			want: "",
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
