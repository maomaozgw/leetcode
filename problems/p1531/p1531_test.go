package p1531

import "testing"

func Test_getLengthOfOptimalCompression(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "aaabcccd",
				k: 2,
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				s: "aabbaa",
				k: 2,
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				s: "aaaaaaaaaaa",
				k: 0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLengthOfOptimalCompression(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("getLengthOfOptimalCompression() = %v, want %v", got, tt.want)
			}
		})
	}
}
