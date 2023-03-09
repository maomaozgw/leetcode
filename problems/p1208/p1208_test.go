package p1208

import "testing"

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s:       "abcd",
				t:       "bcdf",
				maxCost: 3,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				s:       "abcd",
				t:       "cdef",
				maxCost: 3,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				s:       "abcd",
				t:       "acde",
				maxCost: 0,
			},
			want: 1,
		},
		{
			name: "WA 1",
			args: args{
				s:       "ujteygggjwxnfl",
				t:       "nstsenrzttikoy",
				maxCost: 43,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
