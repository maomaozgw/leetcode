package p1249

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
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
				s: "lee(t(c)o)de)",
			},
			want: "lee(t(c)o)de",
		},
		{
			name: "Example 2",
			args: args{
				s: "a)b(c)d",
			},
			want: "ab(c)d",
		},
		{
			name: "Example 3",
			args: args{
				s: "))((",
			},
			want: "",
		},
		{
			name: "WA1",
			args: args{
				s: "())()(((",
			},
			want: "()()",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
