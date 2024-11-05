package p2914

import "testing"

func Test_minChanges(t *testing.T) {
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
				s: "1001",
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				s: "10",
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				s: "0000",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minChanges(tt.args.s); got != tt.want {
				t.Errorf("minChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
