package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// LenTerm tells to us len of terminal window
func LenTerm() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	a := strings.Split(string(out), " ")
	b, e := strconv.Atoi(a[1][:len(a[1])-1])
	if e != nil {
		fmt.Printf("%s", e.Error())
		os.Exit(0)
	}
	return b
}

// CheckString func reports to us how many symbols by ascii-art prog goes to term window
func CheckString(str, font string) int {
	if str == "" {
		return 0
	}
	s := exec.Command("./ascii-art", str, font)
	s.Stdin = os.Stdin
	out, e := s.Output()
	if e != nil {
		fmt.Printf("%s", e.Error())
		os.Exit(0)
	}
	return len(out) / 8
}
