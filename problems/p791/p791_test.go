package p791

import "testing"

func Test_customSortString(t *testing.T) {
	type args struct {
		order string
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				order: "cba",
				s:     "abcd",
			},
			want: "dcba",
		},
		{
			name: "Example 2",
			args: args{
				order: "bcafg",
				s:     "abcd",
			},
			want: "dbca",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := customSortString(tt.args.order, tt.args.s); got != tt.want {
				t.Errorf("customSortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
