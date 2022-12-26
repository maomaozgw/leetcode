package p790

import "testing"

func Test_numTilings(t *testing.T) {
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
				n: 3,
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilings(tt.args.n); got != tt.want {
				t.Errorf("numTilings() = %v, want %v", got, tt.want)
			}
		})
	}
}
