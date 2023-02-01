package p1491

import "testing"

func Test_average(t *testing.T) {
	type args struct {
		salary []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example 1",
			args: args{
				salary: []int{4000, 3000, 1000, 2000},
			},
			want: 2500.0,
		},
		{
			name: "Example 2",
			args: args{
				salary: []int{1000, 2000, 3000},
			},
			want: 2000.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args.salary); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
