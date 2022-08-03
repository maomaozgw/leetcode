package p729

import "testing"

func TestMyCalendar_Book(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				start: 47,
				end:   50,
			},
			want: true,
		},
		{
			name: "Example 1.1",
			args: args{
				start: 33,
				end:   41,
			},
			want: true,
		},
		{
			name: "Example 1.2",
			args: args{
				start: 39,
				end:   45,
			},
			want: false,
		},
		{
			name: "Example 1.3",
			args: args{
				start: 33,
				end:   42,
			},
			want: false,
		},
		{
			name: "Example 1.4",
			args: args{
				start: 25,
				end:   32,
			},
			want: true,
		},
		{
			name: "Example 1.5",
			args: args{
				start: 26,
				end:   35,
			},
			want: false,
		},
		{
			name: "Example 1.6",
			args: args{
				start: 19,
				end:   25,
			},
			want: true,
		},
		{
			name: "Example 1.7",
			args: args{
				start: 3,
				end:   8,
			},
			want: true,
		},
		{
			name: "Example 1.8",
			args: args{
				start: 8,
				end:   13,
			},
			want: true,
		},
		{
			name: "Example 1.9",
			args: args{
				start: 18,
				end:   27,
			},
			want: false,
		},
	}
	this := Constructor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := this.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
