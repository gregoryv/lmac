package lmac

import (
	"encoding/hex"
	"strings"
)

// Find registered organization for the given MAC
// Allowed formats are
//
//	00:00:00:00:00:aa
//	00-00-00-00-00-aa
//	0000000000aa
//
// The values are case insensitive.
func Find(mac string) string {
	p, err := lprefix(mac)
	if err != nil {
		return "unknown"
	}

	org := mal[p]
	if org != "" {
		return org
	}
	return "unknown"
}

func lprefix(v string) ([3]byte, error) {
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
