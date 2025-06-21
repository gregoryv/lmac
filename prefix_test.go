package lmac

import "testing"

func Test_prefixL(t *testing.T) {
	cases := map[string][3]byte{
		"00:00:01": {0, 0, 1},
		"ff:00:01": {255, 0, 1},
		"ff0001":   {255, 0, 1},
		"ff-00-01": {255, 0, 1},

		"00:00:01:aa:bb": {0, 0, 1},
	}
	for arg, exp := range cases {
		got, err := prefixL(arg)
		if got != exp {
			t.Errorf("\ngot: %v\nexp: %v\nerr: %v", got, exp, err)
		}
	}
}

func Test_prefixM(t *testing.T) {
	cases := map[string][4]byte{
		"00:00:00:a1": {0, 0, 0, 0xa0},
	}
	for arg, exp := range cases {
		got, err := prefixM(arg)
		if got != exp {
			t.Errorf("\ngot: %v\nexp: %v\nerr: %v", got, exp, err)
		}
	}
}
