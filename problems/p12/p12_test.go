package p12

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "Example 2",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "Exapmle 3",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "WA 1",
			args: args{
				num: 20,
			},
			want: "XX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
