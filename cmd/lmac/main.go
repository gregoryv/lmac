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
		fmt.Fprintln(os.Stdout, "Usage:", os.Args[0], "[OPTIONS] MAC...", `

Either provide a list of MAC strings as arguments or a list
of MAC values on stdin, one on each line.

OPTIONS
  -h, --help	Show this help and exit.`)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stdout, `
Example

  Lookup network devices
  $ arp -n | awk '{print $3 " " $1}' | lmac
  HWaddress Address
  f4:fe:fb:2e:c7:bc 192.168.1.55 Samsung_Electronics_Co.,Ltd
  f0:9f:c2:60:2b:17 192.168.1.213 Ubiquiti
  30:05:5c:a1:2a:77 192.168.1.190 Brother_industries

  ...


  Lookup specific mac
  $ lmac F8:1A:2B:00:00:FA
  F8:1A:2B:00:00:FA Google`)

		fmt.Fprintln(os.Stdout, "\nLast updated", source.LastUpdate)
	}
	raw := flag.Bool("r", false,
		"Use raw organization name",
	)
	flag.Parse()

	// if arguments are given
	for _, mac := range flag.Args() {
		fmt.Println(mac, tidyOrg(lmac.Lookup(mac)))
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
		parts := strings.Split(line, " ")
		mac := parts[0]
		org := lmac.Lookup(mac)
		if !*raw {
			org = tidyOrg(org)
		}
		fmt.Println(line, org)
	}
}

func tidyOrg(org string) string {
	org = strings.TrimSpace(org)
	org = strings.TrimSuffix(org, ".")
	for _, v := range stripSuffixes {
		org = strings.TrimSuffix(org, v)
	}
	org = strings.ReplaceAll(org, " ", "_")
	org = strings.Trim(org, ".,_")
	return org
}

var stripSuffixes = []string{
	" Inc",
	" INC",
	" AB",
	", LTD",
	",Inc",
}
