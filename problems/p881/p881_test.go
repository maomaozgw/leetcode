package p881

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				people: []int{1, 2},
				limit:  3,
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				people: []int{3, 2, 2, 1},
				limit:  3,
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				people: []int{3, 5, 3, 4},
				limit:  5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
