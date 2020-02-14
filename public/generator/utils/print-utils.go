package utils

import (
	"fmt"
	"os"
	"strings"
)

// PrintToTerminal func apply all parsed flags whithin generating ASCII art
func PrintToTerminal(p *Params) {
	//fmt.Println(p.StrInput)
	newStr := ""
	for _, v := range p.StrInput {
		if v < 127 {
			newStr = newStr + string(v)
		}
	}
	p.StrInput = newStr
	arr := strings.Split(p.StrInput, "\r\n")
	//lH -- Hight of letter
	asciiFile, lH, err := ChooseFont(p)
	if err != "" {
		fmt.Println(err)
		os.Exit(0)
	}

	//fmt.Printf("%v\n", p.CharArr)
	// //fmt.Println(p.ColorOfSlice)
	// fmt.Printf("%v%s\n", p.IntArr, p.ColorOfInterval)
	// fmt.Printf("%s    :  %s\n", p.Slice, p.ColorOfSlice)

	for i := range arr {
		text, interval, sep := FillIntervals(arr[i], p)
		index := 0
		//spec variables for using colorSlice
		indexOfSlice := 0
		sliceBool := false
		//^^^^^
		for h := 0; h < lH; h++ {
			//use color settings
			if p.Color {
				fmt.Printf("%v%v", SRG(p.InitColorState), SRG(p.InitBgColorState))
			}
			fmt.Printf("%s", interval[index])
			index++
			for j := range text {
				for k := range text[j] {
					if p.Interval {
						ColorInterval(k, p.IntArr, p.ColorOfInterval)
					}
					if p.SliceBool {
						ColorSlice(text[j], p.Slice, p.ColorOfSlice, k, &indexOfSlice, &sliceBool)
					}

					if p.Char {
						ColorChar(rune(text[j][k]), p.CharArr)
					}
					fmt.Printf("%s", (asciiFile[((int(text[j][k])-32)*8)+h]))
					if p.SliceBool || p.Interval || p.Char {
						fmt.Printf("%v", SRG("0"))
						fmt.Printf("%v%v", SRG(p.InitColorState), SRG(p.InitBgColorState))
					}
				}
				if len(text) == 1 {
					index++
				}
				if j != len(text)-1 {
					if sep {
						fmt.Printf("%s", (asciiFile[h]))
					}
					fmt.Printf("%s", interval[index])
					index++
				}
			}
			fmt.Printf("%s", interval[index])
			index = 0
			if p.Color {
				fmt.Printf("%v", SRG("0"))
			}
			fmt.Println()
		}

	}
}

// PrintToFile func using when flag <--output=> activated
func PrintToFile(p *Params) {
	arr := strings.Split(p.StrInput, "\\n")
	//lH -- Hight of letter
	arrString, lH, err := ChooseFont(p)
	if err != "" {
		fmt.Println(err)
		os.Exit(0)
	}
	file2, e := os.Create(p.FileToOutput)
	if e != nil {
		fmt.Printf("%s", e.Error())
		fmt.Println()
		os.Exit(0)
	}
	for i := range arr {
		for k := 0; k < lH; k++ {
			for j := range arr[i] {
				fmt.Fprintf(file2, "%s", (arrString[((int(arr[i][j])-32)*8)+k]))
			}
			fmt.Fprintln(file2)
		}
	}
}
