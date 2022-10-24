package p1239

import "testing"

func Test_maxLength(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				arr: []string{"un", "iq", "ue"},
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				arr: []string{"cha", "r", "act", "ers", "abcde"},
			},
			want: 6,
		},
		{
			name: "Example 3",
			args: args{
				arr: []string{"abcdefghijklmnopqrstuvwxyz"},
			},
			want: 26,
		},
		{
			name: "WA 1",
			args: args{
				arr: []string{"aa", "bb"},
			},
			want: 0,
		},
		{
			name: "UC 1",
			args: args{
				arr: []string{"aaa", "a"},
			},
			want: 1,
		},
		{
			name: "WA 4",
			args: args{
				arr: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLength(tt.args.arr); got != tt.want {
				t.Errorf("maxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
