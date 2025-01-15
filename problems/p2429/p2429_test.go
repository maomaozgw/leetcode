package p2429

import "testing"

func Test_minimizeXor(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				num1: 3,
				num2: 5,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				num1: 1,
				num2: 12,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeXor(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("minimizeXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
