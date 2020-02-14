package utils

import "strings"

// CheckStack func cheking stack of strings with tempory stacks constucted with different fonts
func CheckStack(st []string, p *Params) (rune, bool) {
	tempStack := make([]string, 8)
	arrFonts := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}
	for i := range arrFonts {
		p.Font = arrFonts[i]
		font, lH, _ := ChooseFont(p)
		for k := ' '; k <= '~'; k++ {
			for j := 0; j < lH; j++ {
				tempStack[j] = font[(int(k)-32)*lH+j]
			}
			if CompareStack(st, tempStack) {
				return k, true
			}
		}
	}
	return '_', false
}

// CompareStack func using to compate tempory stack and stack from --reverse file
func CompareStack(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if strings.Compare(arr1[i], arr2[i]) != 0 {
			return false
		}
	}
	return true
}
