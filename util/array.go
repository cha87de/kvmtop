package util

// RemoveFromArray removes the element with key `r` from array `s`
func RemoveFromArray(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
