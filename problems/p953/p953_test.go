package p953

import "testing"

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				words: []string{
					"hello", "leetcode",
				},
				order: "hlabcdefgijkmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				words: []string{
					"word", "world", "row",
				},
				order: "worldabcefghijkmnpqstuvxyz",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				words: []string{
					"apple", "app",
				},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: false,
		},
		{
			name: "UE 1",
			args: args{
				words: []string{
					"app", "apple",
				},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			want: true,
		},
		{
			name: "WA 1",
			args: args{
				words: []string{
					"kuvp", "q",
				},
				order: "ngxlkthsjuoqcpavbfdermiywz",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
