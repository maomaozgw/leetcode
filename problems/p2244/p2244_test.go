package p2244

import "testing"

func Test_minimumRounds(t *testing.T) {
	type args struct {
		tasks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				tasks: []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4},
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				tasks: []int{2, 3, 3},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRounds(tt.args.tasks); got != tt.want {
				t.Errorf("minimumRounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
