package p443

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				chars: []byte{
					'a', 'a', 'b', 'b', 'c', 'c', 'c',
				},
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				chars: []byte{'a'},
			},
			want: 1,
		},
		{
			name: "UE 1",
			args: args{
				chars: []byte{'a', 'b', 'c'},
			},
			want: 3,
		},
		{
			name: "UE 2",
			args: args{
				chars: []byte{'a', 'a', 'a'},
			},
			want: 2,
		},
		{
			name: "WA 1",
			args: args{
				chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
