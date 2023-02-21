package p709

import "testing"

func Test_toLowerCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s: "Hello",
			},
			want: "hello",
		},
		{
			name: "Example 2",
			args: args{
				s: "here",
			},
			want: "here",
		},
		{
			name: "Example 3",
			args: args{
				s: "LOVELY",
			},
			want: "lovely",
		},
		{
			name: "WA 1",
			args: args{
				s: "al&phaBET",
			},
			want: "al&phabet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.s); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
