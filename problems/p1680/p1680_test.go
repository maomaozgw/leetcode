package p1680

import "testing"

func Test_concatenatedBinary(t *testing.T) {
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
				n: 1,
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				n: 3,
			},
			want: 27,
		},
		{
			name: "Example 3",
			args: args{
				n: 12,
			},
			want: 505379714,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatenatedBinary(tt.args.n); got != tt.want {
				t.Errorf("concatenatedBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
