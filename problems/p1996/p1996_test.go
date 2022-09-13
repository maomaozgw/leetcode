package p1996

import "testing"

func Test_numberOfWeakCharacters(t *testing.T) {
	type args struct {
		properties [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				properties: [][]int{{5, 5}, {6, 3}, {3, 6}},
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				properties: [][]int{{2, 2}, {3, 3}},
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				properties: [][]int{
					{1, 5}, {10, 4}, {4, 3},
				},
			},
			want: 1,
		},
		{
			name: "WA 1",
			args: args{
				properties: [][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWeakCharacters(tt.args.properties); got != tt.want {
				t.Errorf("numberOfWeakCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
