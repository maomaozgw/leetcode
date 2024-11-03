package p796

import "testing"

func Test_rotateString(t *testing.T) {
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
				s:    "abcde",
				goal: "cdeab",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s:    "abcde",
				goal: "abced",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
