package p223

import "testing"

func Test_computeArea(t *testing.T) {
	type args struct {
		ax1 int
		ay1 int
		ax2 int
		ay2 int
		bx1 int
		by1 int
		bx2 int
		by2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				ax1: -3,
				ay1: 0,
				ax2: 3,
				ay2: 4,
				bx1: 0,
				by1: -1,
				bx2: 9,
				by2: 2,
			},
			want: 45,
		},
		{
			name: "Example 2",
			args: args{
				ax1: -2,
				ay1: -2,
				ax2: 2,
				ay2: 2,
				bx1: -2,
				by1: -2,
				bx2: 2,
				by2: 2,
			},
			want: 16,
		},
		{
			name: "WA 1",
			args: args{
				ax1: 0,
				ay1: 0,
				ax2: 0,
				ay2: 0,
				bx1: -1,
				by1: -1,
				bx2: 1,
				by2: 1,
			},
			want: 4,
		},
		{
			name: "No overlap X",
			args: args{
				ax1: -2,
				ay1: 0,
				ax2: -1,
				ay2: 1,
				bx1: 0,
				by1: 0,
				bx2: 1,
				by2: 1,
			},
			want: 2,
		},
		{
			name: "No overlap Y",
			args: args{
				ax1: 0,
				ay1: 1,
				ax2: 1,
				ay2: 2,
				bx1: 0,
				by1: -1,
				bx2: 1,
				by2: 0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeArea(tt.args.ax1, tt.args.ay1, tt.args.ax2, tt.args.ay2, tt.args.bx1, tt.args.by1, tt.args.bx2, tt.args.by2); got != tt.want {
				t.Errorf("computeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
