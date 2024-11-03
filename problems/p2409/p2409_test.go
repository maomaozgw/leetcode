package p2409

import "testing"

func Test_isCircularSentence(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				sentence: "leetcode exercises sound delightful",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				sentence: "eetcode",
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				sentence: "Leetcode is cool",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCircularSentence(tt.args.sentence); got != tt.want {
				t.Errorf("isCircularSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
