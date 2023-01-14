package p278

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n   int
		bad int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:   5,
				bad: 4,
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				n:   1,
				bad: 1,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				n:   5,
				bad: 1,
			},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{
				n:   5,
				bad: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bad = tt.args.bad
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
