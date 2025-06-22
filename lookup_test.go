package lmac

import (
	"fmt"
)

func ExampleLookup() {
	// ma-l           ........
	fmt.Println(Lookup("00:03:08:00:00:00"))
	fmt.Println(Lookup("00177a000000"))
	fmt.Println()
	// ma-m           ..........
	fmt.Println(Lookup("00:55:da:f0:01:00"))
	fmt.Println(Lookup("04:58:5D:89:99:99"))
	fmt.Println()
	// ma-s           .............
	fmt.Println(Lookup("00:1b:c5:04:b1:11"))
	fmt.Println()
	// unrecognized
	fmt.Println(Lookup("jibberish"))
	fmt.Println(Lookup("ff:ff:ff:ff:ff"))
	// output:
	// AM Communications, Inc.
	// ASSA ABLOY AB
	//
	// Private
	// JRK VISION
	//
	// Silicon Controls
}
