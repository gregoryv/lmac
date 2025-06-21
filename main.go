package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
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
		fmt.Fprintln(os.Stderr, "missing data on stdin")
		os.Exit(0)
	}
	// sort keys
	keys := make([]string, 0, len(oui))
	for k := range oui {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	// write go file
	fmt.Print(`
// GENERATED, DO NOT EDIT!

package main

var result = map[string]string{
`)
	for _, k := range keys {
		fmt.Printf("\t%q: %q,\n", k, oui[k])
	}
	fmt.Println("}")
}

// Parse http://standards-oui.ieee.org/oui/oui.csv to the given map
func Parse(oui map[string]string, r io.Reader) {
	var (
		done    = make(chan struct{})
		parsing = make(chan struct{})
		once    sync.Once
	)

	go func() {
		defer close(done)
		in := csv.NewReader(r)
		for {
			record, err := in.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				log.Print(err)
				continue
			}
			// if scanner is blocked forever, outside knows this
			once.Do(func() { close(parsing) })
			if len(record) < 3 {
				continue
			}
			mac := record[1]
			org := record[2]
			key := strings.ToLower(mac)
			oui[key] = org
		}
	}()

	select {
	case <-time.After(time.Millisecond):
		// if nothing is read on the reader, ie. nothing on stdin
	case <-parsing:
		// scan has started
		<-done
	}
}
