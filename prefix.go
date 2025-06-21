package lmac

import (
	"encoding/hex"
	"strings"
)

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
