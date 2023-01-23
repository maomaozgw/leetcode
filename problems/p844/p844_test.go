package p844

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				s: "a#c",
				t: "b",
			},
			want: false,
		},
		{
			name: "WA 1",
			args: args{
				s: "bxj##tw",
				t: "bxo#j##tw",
			},
			want: true,
		},
		{
			name: "WA 2",
			args: args{
				s: "bxj##tw",
				t: "bxj###tw",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
