// Command lmac looks up MAC addresses against registered
// organizations.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gregoryv/lmac"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, "Usage:", os.Args[0], "MAC...", `

Either provide a list of MAC strings as arguments or a list
of MAC values on stdin, one on each line.
`)
		flag.PrintDefaults()
	}
	flag.Parse()

	for _, mac := range flag.Args() {
		fmt.Println(mac, lmac.Find(mac))
	}

	var (
		done     = make(chan struct{})
		scanning = make(chan struct{})
		once     sync.Once
	)
	go func() {
		defer close(done)
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			once.Do(func() { close(scanning) })
			line := strings.TrimSpace(s.Text())
			if len(line) == 0 {
				continue
			}
			i := strings.Index(line, " ")
			mac := line
			if i > 0 {
				mac = line[:i]
			}

			fmt.Println(line, lmac.Find(mac))
		}
	}()

	select {
	case <-time.After(5 * time.Millisecond):
	case <-scanning:
		<-done
	}
}
