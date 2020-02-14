package utils

import "strings"

// SRG func aplies srg params in order
func SRG(params ...string) string {
	var sb strings.Builder
	sb.Grow(len(params)*4 + 2)
	sb.WriteString(CSI)
	for i, v := range params {
		if i != 0 {
			sb.WriteRune(';')
		}
		sb.WriteString(string(v))
	}
	sb.WriteRune('m')
	return sb.String()
}
