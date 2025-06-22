package lmac

import (
	"encoding/hex"
	"strings"
)

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
