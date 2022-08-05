package p899

import "testing"

func Test_orderlyQueue(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s: "cba",
				k: 1,
			},
			want: "acb",
		},
		{
			name: "Example 2",
			args: args{
				s: "baaca",
				k: 3,
			},
			want: "aaabc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderlyQueue(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("orderlyQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
