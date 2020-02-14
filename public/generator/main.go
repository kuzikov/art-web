package main

import (
	"os"

	u "./utils"
)

func main() {
	///default params///
	p := u.Params{
		InitColorState:   u.FgReset,
		InitBgColorState: u.BgReset,
		Color:            false,
		Output:           false,
		Align:            false,
		Reverse:          false,
		IsCorrectInput:   false,
		Pipe:             false,
		Font:             "standard.txt"}
	defer os.Exit(0)
	u.CheckArgs(os.Args[1:], &p)
	if p.Reverse {
		u.ReadFromFile(&p)
	} else if p.Output {
		u.PrintToFile(&p)
	} else {
		u.PrintToTerminal(&p)
	}
}
