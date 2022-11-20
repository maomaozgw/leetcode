package p224

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "1 + 1",
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				s: " 2-1 + 2 ",
			},
			want: 3,
		},

		{
			name: "Example 3",
			args: args{
				s: "(1+(4+5+2)-3)+(6+8)",
			},
			want: 23,
		},
		{
			name: "Wa 1",
			args: args{
				s: "(1)",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
