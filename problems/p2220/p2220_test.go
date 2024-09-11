package p2220

import "testing"

func Test_minBitFlips(t *testing.T) {
	type args struct {
		start int
		goal  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				start: 10,
				goal:  7,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				start: 3,
				goal:  4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBitFlips(tt.args.start, tt.args.goal); got != tt.want {
				t.Errorf("minBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
