//go:build ignore

// Command to generate Go maps of MAC to organization/company
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
	genFile("mal", "ma-l.go", "source/ma-l.csv")
	genFile("mam", "ma-m.go", "source/ma-m.csv")
	genFile("mas", "ma-s.go", "source/ma-s.csv")
}

func genFile(varname, out, src string) {
	oui := make(map[string]string)
	fh, _ := os.Open(src)
	defer fh.Close()
	Parse(oui, fh)

	if len(oui) == 0 {
		fmt.Fprintln(os.Stderr, "missing data")
		os.Exit(0)
	}
	// sort keys
	keys := make([]string, 0, len(oui))
	for k := range oui {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	// write go file
	w, _ := os.Create(out)
	fmt.Fprintf(w, `
// GENERATED, DO NOT EDIT!

package lmac

var %s = map[string]string{
`, varname)
	for _, k := range keys {
		fmt.Fprintf(w, "\t%q: %q,\n", k, oui[k])
	}
	fmt.Fprintln(w, "}")
	w.Close()
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
			if org == "Organization Name" {
				continue
			}
			tmp := strings.ToLower(mac)
			var key string
			for i := 0; i < len(tmp); i += 2 {
				if i > 0 {
					key += ":"
				}
				if i+2 > len(tmp) {
					key += tmp[i:]
					continue
				}
				key += tmp[i : i+2]
			}
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
