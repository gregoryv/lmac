package lmac

// Lookup registered organization for the given MAC
// Allowed formats are
//
//	00:00:00:00:00:aa
//	00-00-00-00-00-aa
//	0000000000aa
//
// The values are case insensitive.
func Lookup(mac string) string {
	if res := findMas(mac); res != "" {
		return res
	}

	if res := findMam(mac); res != "" {
		return res
	}

	return findMal(mac)
}

func findMal(mac string) string {
	var key [3]byte
	if err := prefix(key[:], mac); err != nil {
		return ""
	}
	return mal[key]
}

func findMam(mac string) string {
	var key [4]byte
	if err := prefix(key[:], mac); err != nil {
		return ""
	}
	key[3] = key[3] & 0xf0
	return mam[key]
}

func findMas(mac string) string {
	var key [5]byte
	if err := prefix(key[:], mac); err != nil {
		return ""
	}
	key[4] = key[4] & 0xf0
	return mas[key]
}
