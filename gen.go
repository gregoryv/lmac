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
	genMAL("mal", "mal.go", "source/mal.csv")
	genMAM("mam", "mam.go", "source/mam.csv")
	genMAS("mas", "mas.go", "source/mas.csv")
}

func genMAL(varname, out, src string) {
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

var %s = map[[3]byte]string{
`, varname)
	for _, mac := range keys {
		p, err := prefixL(mac)
		if err != nil {
			log.Fatal(err)
		}
		key := fmt.Sprintf("[3]byte{%v, %v, %v}", p[0], p[1], p[2])
		fmt.Fprintf(w, "\t%s: %q,\n", key, oui[mac])
	}
	fmt.Fprintln(w, "}")
	w.Close()
}

func genMAM(varname, out, src string) {
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

var %s = map[[4]byte]string{
`, varname)
	for _, mac := range keys {
		p, err := prefixM(mac + "0")
		if err != nil {
			log.Fatal(err)
		}
		key := fmt.Sprintf("[4]byte{%v, %v, %v, %v}", p[0], p[1], p[2], p[3])
		fmt.Fprintf(w, "\t%s: %q,\n", key, oui[mac])
	}
	fmt.Fprintln(w, "}")
	w.Close()
}

func genMAS(varname, out, src string) {
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

var %s = map[[5]byte]string{
`, varname)
	for _, mac := range keys {
		p, err := prefixS(mac + "0")
		if err != nil {
			log.Fatal(err)
		}
		key := fmt.Sprintf("[5]byte{%v, %v, %v, %v, %v}", p[0], p[1], p[2], p[3], p[4])
		fmt.Fprintf(w, "\t%s: %q,\n", key, oui[mac])
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

// ----------------------------------------
// same as prefix.go

func prefixL(v string) ([3]byte, error) {
	v = strings.ReplaceAll(v, ":", "")
	v = strings.ReplaceAll(v, "-", "")

	var p [3]byte
	raw, err := hex.DecodeString(v)
	if err != nil {
		return p, err
	}

	p[0] = raw[0]
	p[1] = raw[1]
	p[2] = raw[2]
	return p, nil
}

func prefixM(v string) ([4]byte, error) {
	v = strings.ReplaceAll(v, ":", "")
	v = strings.ReplaceAll(v, "-", "")

	var p [4]byte
	raw, err := hex.DecodeString(v)
	if err != nil {
		return p, err
	}

	p[0] = raw[0]
	p[1] = raw[1]
	p[2] = raw[2]
	p[3] = raw[3] & 0xf0
	return p, nil
}

func prefixS(v string) ([5]byte, error) {
	v = strings.ReplaceAll(v, ":", "")
	v = strings.ReplaceAll(v, "-", "")

	var p [5]byte
	raw, err := hex.DecodeString(v)
	if err != nil {
		return p, err
	}

	p[0] = raw[0]
	p[1] = raw[1]
	p[2] = raw[2]
	p[3] = raw[3]
	p[4] = raw[4] & 0xf0
	return p, nil
}
