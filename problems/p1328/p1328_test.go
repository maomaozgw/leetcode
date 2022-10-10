package p1328

import "testing"

func Test_breakPalindrome(t *testing.T) {
	type args struct {
		palindrome string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				palindrome: "abccba",
			},
			want: "aaccba",
		},
		{
			name: "Example 2",
			args: args{
				palindrome: "a",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakPalindrome(tt.args.palindrome); got != tt.want {
				t.Errorf("breakPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
