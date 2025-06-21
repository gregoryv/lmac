package lmac

import (
	"fmt"
	"testing"
)

func ExampleFind() {
	// ma-l           ........
	fmt.Println(Find("00:03:08:00:00:00"))
	fmt.Println(Find("00177a000000"))
	fmt.Println()
	// ma-m           ..........
	fmt.Println(Find("00:55:da:f0:01:00"))
	fmt.Println(Find("04:58:5D:89:99:99"))
	fmt.Println()
	// ma-s           .............
	fmt.Println(Find("00:1b:c5:04:b1:11"))
	// output:
	// AM Communications, Inc.
	// ASSA ABLOY AB
	//
	// Private
	// JRK VISION
	//
	// Silicon Controls
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
		{
			arg: "ff:ff:ff:ff:ff",
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
