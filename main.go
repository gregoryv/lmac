package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

func main() {
	oui := make(map[string]string)
	Parse(oui, os.Stdin)

	if len(oui) == 0 {
		fmt.Fprintln(os.Stderr, "missing data, you must pipe oui.txt on stdin")
		os.Exit(0)
	}
	// sort keys
	keys := make([]string, 0, len(oui))
	for k := range oui {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	// write go file
	fmt.Print(`package main

// http://standards-oui.ieee.org/oui/oui.txt
var oui = map[string]string{
`)
	for _, k := range keys {
		fmt.Printf("\t%q: %q,\n", k, oui[k])
	}
	fmt.Println("}")
}

// Parse http://standards-oui.ieee.org/oui/oui.txt to the given map
func Parse(oui map[string]string, r io.Reader) {
	var (
		done    = make(chan struct{})
		parsing = make(chan struct{})
		once    sync.Once
	)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// if scanner is blocked forever, outside knows this
			once.Do(func() { close(parsing) })

			line := s.Bytes()
			if len(line) < 3 || line[2] != '-' {
				continue
			}
			i := bytes.Index(line, []byte{' '})
			mac := line[:i]
			mac = bytes.Replace(mac, []byte{'-'}, []byte{':'}, 2)

			j := bytes.Index(line, []byte{')'})
			org := bytes.TrimSpace(line[j+1:])

			key := string(bytes.ToLower(mac))
			// normalize organization names
			val := string(bytes.ToLower(org))
			val = strings.TrimSuffix(val, ".")
			val = strings.TrimSuffix(val, "ltd")
			val = strings.TrimSuffix(val, "inc")
			val = strings.TrimSuffix(val, "corp")
			val = strings.TrimSpace(val)
			val = strings.TrimSuffix(val, ",")
			val = strings.TrimSuffix(val, ".")
			val = strings.TrimSuffix(val, "co")
			val = strings.TrimSpace(val)
			val = strings.Title(val)
			oui[key] = val
		}
		close(done)
	}()

	select {
	case <-time.After(time.Millisecond):
		// if nothing is read on the reader, ie. nothing on stdin
	case <-parsing:
		// scan has started
		<-done
	}
}
