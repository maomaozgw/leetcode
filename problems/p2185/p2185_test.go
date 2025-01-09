package p2185

import "testing"

func Test_prefixCount(t *testing.T) {
	type args struct {
		words []string
		pref  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				words: []string{"pay", "attention", "practice", "attend"},
				pref:  "at",
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				words: []string{"leetcode", "win", "loops", "success"},
				pref:  "code",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixCount(tt.args.words, tt.args.pref); got != tt.want {
				t.Errorf("prefixCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
