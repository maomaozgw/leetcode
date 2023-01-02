package p1823

import "testing"

func Test_findTheWinner(t *testing.T) {
	type args struct {
		n int
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
				n: 5,
				k: 2,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				n: 6,
				k: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheWinner(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
