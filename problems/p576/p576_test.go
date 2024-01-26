package p576

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				m:           2,
				n:           2,
				maxMove:     2,
				startRow:    0,
				startColumn: 0,
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				m:           1,
				n:           3,
				maxMove:     3,
				startRow:    0,
				startColumn: 1,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); got != tt.want {
				t.Errorf("findPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
