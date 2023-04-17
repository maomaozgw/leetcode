package p441

import "testing"

func Test_arrangeCoins(t *testing.T) {
	type args struct {
		n int
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
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				n: 8,
			},
			want: 3,
		},
		{
			name: "UE 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "UE 2",
			args: args{
				n: 6,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoins(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
