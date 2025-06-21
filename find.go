package lmac

func Find(v string) string {
	org := mal[v]
	if org != "" {
		return org
	}
	return "unknown"
}
