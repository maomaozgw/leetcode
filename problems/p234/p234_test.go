package p234

import (
	"testing"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				head: listnode.NewG(1, 2, 2, 1),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				head: listnode.NewG(1, 2),
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				head: listnode.NewG(1, 2, 1),
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				head: listnode.NewG(1, 2, 3),
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				head: listnode.NewG(1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
