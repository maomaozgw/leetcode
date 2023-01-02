package p520

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				word: "USAZ",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				word: "flagaz",
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				word: "FlaG",
			},
			want: false,
		},
		{
			name: "Example 4",
			args: args{
				word: "Flag",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.args.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
