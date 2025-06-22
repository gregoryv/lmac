//go:build ignore

// Command to generate Go maps of MAC to organization/company
package main

import (
	"encoding/csv"
	"encoding/hex"
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
	gen(3, "", "mal", "mal.go", "source/mal.csv")
	gen(4, "0", "mam", "mam.go", "source/mam.csv")
	gen(5, "0", "mas", "mas.go", "source/mas.csv")
}

func gen(size int, last, varname, out, src string) {
	oui := make(map[string]string)
	fh, _ := os.Open(src)
	defer fh.Close()

	keys := Parse(oui, fh)

	// write go file
	w, _ := os.Create(out)
	fmt.Fprintf(w, header, varname, size)
	for _, mac := range keys {
		p := prefix(size, mac+last)
		fmt.Fprintln(w, mapline(p, oui[mac]))
	}
	fmt.Fprintln(w, "}")
	w.Close()
}

const header = `
// GENERATED, DO NOT EDIT!

package lmac

var %s = map[[%v]byte]string{
`

func mapline(p []byte, val string) string {
	format := fmt.Sprintf("[%v]byte{", len(p))
	var args []any
	for i, b := range p {
		args = append(args, b)
		if i == 0 {
			format += "%v"
			continue
		}
		format += ", %v"
	}
	format += "}"
	key := fmt.Sprintf(format, args...)
	return fmt.Sprintf("\t%s: %q,", key, val)
}

// Parse http://standards-oui.ieee.org/oui/oui.csv to the given map
func Parse(oui map[string]string, r io.Reader) []string {
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
	return keys
}

// ----------------------------------------
// similar to the one in prefix.go

func prefix(size int, v string) []byte {
	v = strings.ReplaceAll(v, ":", "")
	v = strings.ReplaceAll(v, "-", "")

	raw, err := hex.DecodeString(v)
	if err != nil {
		log.Fatal(err)
	}
	dst := make([]byte, size)
	copy(dst, raw)
	return dst
}
