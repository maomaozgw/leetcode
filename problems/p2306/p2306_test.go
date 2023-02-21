package p2306

import "testing"

func Test_distinctNames(t *testing.T) {
	type args struct {
		ideas []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				ideas: []string{
					"coffee", "donuts", "time", "toffee",
				},
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				ideas: []string{
					"lack", "back",
				},
			},
			want: 0,
		},
		{
			name: "WA 1",
			args: args{
				ideas: []string{
					"ps", "yn", "eciuehjexw", "pdesv", "e", "dz", "xxdnk", "vdqjhncfj",
				},
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctNames(tt.args.ideas); got != tt.want {
				t.Errorf("distinctNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
