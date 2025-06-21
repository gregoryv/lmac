package lmac

import (
	"encoding/hex"
	"strings"
)

func prefixL(v string) ([3]byte, error) {
	var p [3]byte
	if err := prefix(p[:], v); err != nil {
		return p, err
	}
	return p, nil
}

func prefixM(v string) ([4]byte, error) {
	var p [4]byte
	if err := prefix(p[:], v); err != nil {
		return p, err
	}
	p[3] = p[3] & 0xf0
	return p, nil
}

func prefixS(v string) ([5]byte, error) {
	var p [5]byte
	if err := prefix(p[:], v); err != nil {
		return p, err
	}
	p[4] = p[4] & 0xf0
	return p, nil
}

func prefix(dst []byte, v string) error {
	v = strings.ReplaceAll(v, ":", "")
	v = strings.ReplaceAll(v, "-", "")

	raw, err := hex.DecodeString(v)
	if err != nil {
		return err
	}
	copy(dst, raw)
	return nil
}
