package p838

import "testing"

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				dominoes: "RR.L",
			},
			want: "RR.L",
		},
		{
			name: "Example 2",
			args: args{
				dominoes: ".L.R...LR..L..",
			},
			want: "LL.RR.LLRRLL..",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushDominoes(tt.args.dominoes); got != tt.want {
				t.Errorf("pushDominoes() %s = %v, want %v", tt.args.dominoes, got, tt.want)
			}
		})
	}
}
