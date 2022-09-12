package p948

import "testing"

func Test_bagOfTokensScore(t *testing.T) {
	type args struct {
		tokens []int
		power  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				tokens: []int{100},
				power:  50,
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				tokens: []int{100, 200},
				power:  150,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				tokens: []int{
					100, 200, 300, 400,
				},
				power: 200,
			},
			want: 2,
		},
		{
			name: "WA 1",
			args: args{
				tokens: []int{71, 55, 82},
				power:  54,
			},
			want: 0,
		},
		{
			name: "empty",
			args: args{
				tokens: []int{},
				power:  100,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bagOfTokensScore(tt.args.tokens, tt.args.power); got != tt.want {
				t.Errorf("bagOfTokensScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
