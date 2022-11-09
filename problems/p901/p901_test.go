package p901

import (
	"testing"
)

func TestStockSpanner_Next(t *testing.T) {

	type args struct {
		price int
	}
	this := Constructor()
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1.1",
			args: args{
				price: 100,
			},
			want: 1,
		},
		{
			name: "Example 1.2",
			args: args{
				price: 80,
			},
			want: 1,
		},
		{
			name: "Example 1.3",
			args: args{
				price: 60,
			},
			want: 1,
		},
		{
			name: "Example 1.4",
			args: args{
				price: 70,
			},
			want: 2,
		},
		{
			name: "Example 1.5",
			args: args{
				price: 60,
			},
			want: 1,
		}, {
			name: "Example 1.6",
			args: args{
				price: 75,
			},
			want: 4,
		},
		{
			name: "Example 1.7",
			args: args{
				price: 85,
			},
			want: 6,
		},
		{
			name: "Example 1.8",
			args: args{
				price: 85,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := this.Next(tt.args.price); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
