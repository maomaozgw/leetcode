package p859

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s:    "ab",
				goal: "ba",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s:    "ab",
				goal: "ab",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				s:    "aa",
				goal: "aa",
			},
			want: true,
		},
		{
			name: "WA 1",
			args: args{
				s:    "abcaa",
				goal: "abcbb",
			},
			want: false,
		},
		{
			name: "WA 2",
			args: args{
				s:    "abcd",
				goal: "badc",
			},
			want: false,
		},
		{
			name: "WA 3",
			args: args{
				s:    "abac",
				goal: "abad",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
