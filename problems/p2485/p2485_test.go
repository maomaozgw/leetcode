package p2485

import "testing"

func Test_pivotInteger(t *testing.T) {
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
				n: 8,
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				n: 4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotInteger(tt.args.n); got != tt.want {
				t.Errorf("pivotInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
