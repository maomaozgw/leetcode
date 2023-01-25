package p191

import (
	"strconv"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	binary2Number := func(s string) uint32 {
		num, err := strconv.ParseUint(s, 2, 32)
		if err != nil {
			panic(err)

		}
		return uint32(num)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				num: binary2Number("00000000000000000000000000001011"),
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				num: binary2Number("00000000000000000000000010000000"),
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				num: binary2Number("11111111111111111111111111111101"),
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
