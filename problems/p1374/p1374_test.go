package p1374

import (
	"fmt"
	"testing"
)

func Test_generateTheString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				n: 4,
			},
			want: "pppz",
		},
		{
			name: "Example 2",
			args: args{
				n: 2,
			},
			want: "xy",
		},
		{
			name: "Example 3",
			args: args{
				n: 7,
			},
			want: "holasss",
		},
	}

	valid := func(str string, length int) bool {
		result := map[uint8]int{}
		if len(str) != length {
			return false
		}
		for i := 0; i < len(str); i++ {
			result[str[i]]++
		}
		for _, v := range result {
			if v%2 == 0 {
				return false
			}
		}
		return true
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTheString(tt.args.n); !valid(got, tt.args.n) {
				t.Errorf("generateTheString() = %v", got)
			} else {
				fmt.Println(got)
			}
		})
	}
}
