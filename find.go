package lmac

// Find registered organization for the given MAC
// Allowed formats are
//
//	00:00:00:00:00:aa
//	00-00-00-00-00-aa
//	0000000000aa
//
// The values are case insensitive.
func Find(mac string) string {
	if res := findMal(mac); res != "" && res != "IEEE Registration Authority" {
		return res
	}

	if res := findMam(mac); res != "" && res != "IEEE Registration Authority" {
		return res
	}

	return "unknown"
}

func findMal(mac string) string {
	p, err := prefixL(mac)
	if err != nil {
		return ""
	}

	org := mal[p]
	if org != "" {
		return org
	}

	return ""
}

func findMam(mac string) string {
	p, err := prefixM(mac)
	if err != nil {
		return ""
	}

	org := mam[p]
	if org != "" {
		return org
	}

	return ""
}
