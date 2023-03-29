package p1402

import "testing"

func Test_maxSatisfaction(t *testing.T) {
	type args struct {
		satisfaction []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				satisfaction: []int{-1, -8, 0, 5, -9},
			},
			want: 14,
		},
		{
			name: "Example 2",
			args: args{
				satisfaction: []int{4, 3, 2},
			},
			want: 20,
		},
		{
			name: "Example 3",
			args: args{
				satisfaction: []int{-1, -4, -5},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfaction(tt.args.satisfaction); got != tt.want {
				t.Errorf("maxSatisfaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
