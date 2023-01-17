package p380

import (
	"strconv"
	"strings"
	"testing"
)

func intArg(s string) int {
	v, _ := strconv.Atoi(strings.Trim(s, "[]"))
	return v
}

func TestConstructor(t *testing.T) {
	tests := []struct {
		name     string
		funcs    []string
		values   []string
		excepted []string
	}{
		{
			name:     "insert",
			funcs:    []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
			values:   strings.Split("[],[1],[2],[2],[],[1],[2],[]", ","),
			excepted: strings.Split("null,true,false,true,2,true,false,2", ","),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var set RandomizedSet
			for idx, fn := range tt.funcs {
				switch fn {
				case "RandomizedSet":
					set = Constructor()
					continue
				case "insert":
					val := set.Insert(intArg(tt.values[idx]))
					t.Logf("%s (%s) = %v (%s)", fn, tt.values[idx], val, tt.excepted[idx])
					if (tt.excepted[idx] == "true") != val {
						t.Errorf("%s = %v, want %v", fn, val, tt.excepted[idx])
					}

				case "remove":
					val := set.Remove(intArg(tt.values[idx]))
					t.Logf("%s (%s) = %v (%s)", fn, tt.values[idx], val, tt.excepted[idx])
					if (tt.excepted[idx] == "true") != val {
						t.Errorf("%s = %v, want %v", fn, val, tt.excepted[idx])
					}
				case "getRandom":
					val := set.GetRandom()
					_ = intArg(tt.excepted[idx])
					t.Logf("%s (%s) = %v (%s)", fn, tt.values[idx], val, tt.excepted[idx])
				}
			}
		})
	}
}
