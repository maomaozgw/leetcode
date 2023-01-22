package p299

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				secret: "1807",
				guess:  "7810",
			},
			want: "1A3B",
		},
		{
			name: "Example 2",
			args: args{
				secret: "1123",
				guess:  "0111",
			},
			want: "1A1B",
		},
		{
			name: "UE 1",
			args: args{
				secret: "0123",
				guess:  "1111",
			},
			want: "1A0B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
