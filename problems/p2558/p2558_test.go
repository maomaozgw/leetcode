package p2558

import "testing"

func Test_pickGifts(t *testing.T) {
	type args struct {
		gifts []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				gifts: []int{
					25, 64, 9, 4, 100,
				},
				k: 4,
			},
			want: 29,
		},
		{
			name: "Example 2",
			args: args{
				gifts: []int{
					1, 1, 1, 1,
				},
				k: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickGifts(tt.args.gifts, tt.args.k); got != tt.want {
				t.Errorf("pickGifts() = %v, want %v", got, tt.want)
			}
		})
	}
}
