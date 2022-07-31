package p307

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "WA 1",
			fields: fields{
				data: []int{7, 2, 7, 2, 0, 3, 4, 5},
			},
			args: args{
				left:  1,
				right: 7,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.fields.data)
			if got := this.SumRange(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
