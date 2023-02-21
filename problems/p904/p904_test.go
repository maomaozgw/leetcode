package p904

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		fruits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				fruits: []int{1, 2, 1},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				fruits: []int{0, 1, 2, 2},
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				fruits: []int{1, 2, 3, 2, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.args.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
