package lmac

import "testing"

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
