package p227

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
				s: "3+2*2",
			},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{
				s: "3/2",
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				s: "3+5/2",
			},
			want: 5,
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
