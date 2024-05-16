package p3075

import "testing"

func Test_maximumHappinessSum(t *testing.T) {
	type args struct {
		happiness []int
		k         int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				happiness: []int{1, 2, 3},
				k:         2,
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				happiness: []int{1, 1, 1, 1},
				k:         2,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				happiness: []int{2, 3, 4, 5},
				k:         1,
			},
			want: 5,
		},
		{
			name: "WA 1",
			args: args{
				happiness: []int{12, 1, 42},
				k:         3,
			},
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumHappinessSum(tt.args.happiness, tt.args.k); got != tt.want {
				t.Errorf("maximumHappinessSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
