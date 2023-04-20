package p1346

import "testing"

func Test_checkIfExist(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				arr: []int{10, 2, 5, 3},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				arr: []int{3, 1, 7, 11},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist(tt.args.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
