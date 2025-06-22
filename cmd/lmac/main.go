// Command lmac looks up MAC addresses against registered
// organizations.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gregoryv/lmac"
	"github.com/gregoryv/lmac/source"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, "Usage:", os.Args[0], "MAC...", `

Either provide a list of MAC strings as arguments or a list
of MAC values on stdin, one on each line.

Example

  Scan network for devices
  $ arp -a | awk '{print $4 " " $2}' | lmac
  ...

  Lookup specific mac
  $ lmac F8:1A:2B:00:00:FA
  F8:1A:2B:00:00:FA Google, Inc.`)

		fmt.Fprintln(os.Stdout, "\nLast updated", source.LastUpdate)
		flag.PrintDefaults()
	}
	flag.Parse()

	// if arguments are given
	for _, mac := range flag.Args() {
		fmt.Println(mac, lmac.Lookup(mac))
	}
	if len(flag.Args()) > 0 {
		os.Exit(0)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			continue
		}
		i := strings.Index(line, " ")
		mac := line
		if i > 0 {
			mac = line[:i]
		}

		fmt.Println(line, lmac.Lookup(mac))
	}
}
