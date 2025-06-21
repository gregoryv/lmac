package lmac

import (
	"fmt"
	"testing"
)

func ExampleFind() {
	fmt.Println(Find("00:03:08:00:00:00"))
	fmt.Println(Find("00177a000000"))
	// ma-m example
	fmt.Println(Find("00:55:da:f0:01:00"))
	fmt.Println(Find("04:58:5D:89:99:99"))
	// output:
	// AM Communications, Inc.
	// ASSA ABLOY AB
	// Private
	// JRK VISION
}

func TestFind(t *testing.T) {
	cases := []struct {
		arg string
		exp string
	}{
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
