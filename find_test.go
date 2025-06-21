package lmac

import (
	"testing"
)

func TestFind(t *testing.T) {
	cases := []struct {
		arg string
		exp string
	}{
		{
			arg: "00:03:08",
			exp: "AM Communications, Inc.",
		},
		{
			arg: "ff:ff:tt",
			exp: "unknown",
		},
	}
	for _, c := range cases {
		t.Run(c.arg, func(t *testing.T) {
			got := Find(c.arg)
			if got != c.exp {
				t.Errorf("\ngot: %s\nexp: %s", got, c.exp)
			}
		})
	}
}

func Test_lprefix(t *testing.T) {
	cases := map[string][3]byte{
		"00:00:01": {0, 0, 1},
		"ff:00:01": {255, 0, 1},
		"ff0001":   {255, 0, 1},
		"ff-00-01": {255, 0, 1},

		"00:00:01:aa:bb": {0, 0, 1},
	}
	for arg, exp := range cases {
		got, err := lprefix(arg)
		if got != exp {
			t.Errorf("\ngot: %v\nexp: %v\nerr: %v", got, exp, err)
		}
	}
}
