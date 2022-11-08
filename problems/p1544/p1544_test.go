package p1544

import "testing"

func Test_makeGood(t *testing.T) {
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
				s: "leEeetcode",
			},
			want: "leetcode",
		},
		{
			name: "Example 2",
			args: args{
				s: "abBAcC",
			},
			want: "",
		},
		{
			name: "WA1",
			args: args{
				s: "Pp",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGood(tt.args.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
