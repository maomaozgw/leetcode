package p1700

import "testing"

func Test_countStudents(t *testing.T) {
	type args struct {
		students   []int
		sandwiches []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				students:   []int{1, 1, 0, 0},
				sandwiches: []int{0, 1, 0, 1},
			},
			want: 0,
		},
		{
			name: "Example 1",
			args: args{
				students:   []int{1, 1, 1, 0, 0, 1},
				sandwiches: []int{1, 0, 0, 0, 1, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStudents(tt.args.students, tt.args.sandwiches); got != tt.want {
				t.Errorf("countStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
