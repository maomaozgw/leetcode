package p926

import "testing"

func Test_minFlipsMonoIncr(t *testing.T) {
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
				s: "00110",
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				s: "010110",
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				s: "00011000",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlipsMonoIncr(tt.args.s); got != tt.want {
				t.Errorf("minFlipsMonoIncr() = %v, want %v", got, tt.want)
			}
		})
	}
}
