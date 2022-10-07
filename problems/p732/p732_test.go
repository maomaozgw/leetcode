package p732

import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"testing"
)

func TestMyCalendarThree_Book(t *testing.T) {
	type fields struct {
		pointsValue map[int]int
		h           *binaryheap.Heap
	}
	type args struct {
		start int
		end   int
	}
	c := Constructor()
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1.1",
			args: args{
				start: 10,
				end:   20,
			},
			want: 1,
		},
		{
			name: "Example 1.2",
			args: args{
				50, 60,
			},
			want: 1,
		},
		{
			name: "Example 1.3",
			args: args{
				10, 40,
			},
			want: 2,
		},
		{
			name: "Example 1.4",
			args: args{
				5, 15,
			},
			want: 3,
		},
		{
			name: "Example 1.5",
			args: args{
				5, 10,
			},
			want: 3,
		},
		{
			name: "Example 1.6",
			args: args{
				25, 55,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := c.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
