package p1945

import "testing"

func Test_getLucky(t *testing.T) {
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
				s: "iiii",
				k: 1,
			},
			want: 36,
		},
		{
			name: "Example 2",
			args: args{
				s: "leetcode",
				k: 2,
			},
			want: 6,
		},
		{
			name: "Example 3",
			args: args{
				s: "zbax",
				k: 2,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLucky(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("getLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
