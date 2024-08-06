package p3016

import "testing"

func Test_minimumPushes(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				word: "abcde",
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				word: "xyzxyzxyzxyz",
			},
			want: 12,
		},
		{
			name: "Example 3",
			args: args{
				word: "aabbccddeeffgghhiiiiii",
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPushes(tt.args.word); got != tt.want {
				t.Errorf("minimumPushes() = %v, want %v", got, tt.want)
			}
		})
	}
}
