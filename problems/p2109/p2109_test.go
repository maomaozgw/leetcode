package p2109

import "testing"

func Test_addSpaces(t *testing.T) {
	type args struct {
		s      string
		spaces []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s:      "LeetcodeHelpsMeLearn",
				spaces: []int{8, 13, 15},
			},
			want: "Leetcode Helps Me Learn",
		},
		{
			name: "Example 2",
			args: args{
				s:      "icodeinpython",
				spaces: []int{1, 5, 7, 9},
			},
			want: "i code in py thon",
		},
		{
			name: "Example 3",
			args: args{
				s:      "spacing",
				spaces: []int{0, 1, 2, 3, 4, 5, 6},
			},
			want: " s p a c i n g",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSpaces(tt.args.s, tt.args.spaces); got != tt.want {
				t.Errorf("addSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
