package p2053

import "testing"

func Test_kthDistinct(t *testing.T) {
	type args struct {
		arr []string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				arr: []string{"d", "b", "c", "b", "c", "a"},
				k:   2,
			},
			want: "a",
		},
		{
			name: "Exmaple 2",
			args: args{
				arr: []string{"aaa", "aa", " a"},
				k:   1,
			},
			want: "aaa",
		},
		{
			name: "Exmaple 3",
			args: args{
				arr: []string{"a", "b", "a"},
				k:   3,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthDistinct(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("kthDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
