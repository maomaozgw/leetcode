package p678

import "testing"

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s: "(*)",
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				s: "(*))",
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				s: "**))",
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				s: "*))",
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				s: "(((*))",
			},
			want: true,
		},
		{
			name: "Example 6",
			args: args{
				s: "((((*))",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
