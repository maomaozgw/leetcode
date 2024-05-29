package p1404

import "testing"

func Test_numSteps(t *testing.T) {
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
				s: "1101",
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				s: "10",
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				s: "1",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSteps(tt.args.s); got != tt.want {
				t.Errorf("numSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
