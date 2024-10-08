package p1963

import "testing"

func Test_minSwaps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "][][",
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				s: "]]][[[",
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				s: "[]",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.args.s); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
