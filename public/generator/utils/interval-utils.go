package utils

import (
	"fmt"
	"os"
	"strings"
)

// FillIntervals - intervals using every times when printing to terminal
func FillIntervals(str string, p *Params) ([]string, []string, bool) {
	//fmt.Println("Align")

	if p.Align {
		arr := []string{}
		lenTermWindow := LenTerm()
		///font has at end of name .txt, needed delete this before start func checkString
		lenSymbols := CheckString(str, p.Font[:len(p.Font)-4])
		///the number of pieces of text to be separated by intervals
		array := []string{}
		if lenSymbols == 0 {
			if p.Pipe {
				array = append(array, "|")
				for len(array[0]) < lenTermWindow-1 {
					array[0] += " "
				}
				array = append(array, "")
				array = append(array, "|")
			} else {
				array = append(array, "")
				for len(array[0]) < lenTermWindow {
					array[0] += " "
				}
				array = append(array, "")
				array = append(array, "")
			}
			return []string{""}, array, false
		} else if lenTermWindow < lenSymbols {
			fmt.Println("Terminal window so small for using this option")
			os.Exit(0)
		}
		if p.Textpos == "justify" {
			arr = strings.Split(str, " ")
			if len(arr) == 1 {
				if p.Pipe {
					array = append(array, "|")
					array = append(array, "")
					array = append(array, "|")
					for len(array[2]) <= lenTermWindow-lenSymbols-1 {
						array[2] = " " + array[2]
					}
				} else {
					array = append(array, "")
					array = append(array, "")
					array = append(array, "")
					for len(array[2]) <= lenTermWindow-lenSymbols {
						array[2] = " " + array[2]
					}
				}
				return arr, array, true
			}
		} else {
			arr = append(arr, str)
		}
		var pieces int
		if len(arr) == 1 {
			pieces = 1
		} else {
			pieces = len(arr) - 1
		}

		switch p.Textpos {
		case "left":
			if p.Pipe {
				array = append(array, "|")
				array = append(array, "")
				array = append(array, "|")
				for len(array[2]) <= lenTermWindow-lenSymbols-1 {
					array[2] = " " + array[2]
				}
			} else {
				array = append(array, "")
				array = append(array, "")
				array = append(array, "")
				for len(array[2]) <= lenTermWindow-lenSymbols {
					array[2] = " " + array[2]
				}
			}
			return arr, array, false
		case "right":
			if p.Pipe {
				array = append(array, "|")
				for len(array[0]) <= lenTermWindow-lenSymbols-1 {
					array[0] += " "
				}
				array = append(array, "")
				array = append(array, "|")
			} else {
				array = append(array, "")
				for len(array[0]) <= lenTermWindow-lenSymbols {
					array[0] += " "
				}
				array = append(array, "")
				array = append(array, "")
			}
			return arr, array, false
		case "center":
			if p.Pipe {
				array = append(array, "|")
				array = append(array, "")
				array = append(array, "|")
				for len(array[0])+len(array[2])+lenSymbols < lenTermWindow-1 {
					array[0] += " "
					array[2] = " " + array[2]
				}
				for len(array[0])+len(array[2])+lenSymbols <= lenTermWindow {
					array[2] = " " + array[2]
				}
			} else {
				array = append(array, "")
				array = append(array, "")
				array = append(array, "")
				for len(array[0])+len(array[2])+lenSymbols < lenTermWindow-1 {
					array[0] += " "
					array[2] = " " + array[2]
				}
				for len(array[0])+len(array[2])+lenSymbols <= lenTermWindow {
					array[2] = " " + array[2]
				}
			}
			return arr, array, false
		case "justify":
			intervalsArray := make([]string, pieces)
			if p.Pipe {
				for LenIntervalArray(intervalsArray)+lenSymbols+2 <= lenTermWindow {
					for ind := range intervalsArray {
						if LenIntervalArray(intervalsArray)+lenSymbols+1 == lenTermWindow {
							goto here
						}
						intervalsArray[ind] += " "
					}
				}
			} else {
				for LenIntervalArray(intervalsArray)+lenSymbols <= lenTermWindow {
					for ind := range intervalsArray {
						if LenIntervalArray(intervalsArray)+lenSymbols-1 == lenTermWindow {
							goto here
						}
						intervalsArray[ind] += " "
					}
				}
			}

		here:
			if p.Pipe {
				array = append(array, "|")
				for i := range intervalsArray {
					array = append(array, intervalsArray[i])
				}
				array = append(array, "|")
			} else {
				array = append(array, "")
				for i := range intervalsArray {
					array = append(array, intervalsArray[i])
				}
				array = append(array, "")
			}
			return arr, array, true
		}
	}
	return []string{str}, []string{"", "", ""}, false
}

// LenIntervalArray  needed to know how many symbols in array
func LenIntervalArray(arr []string) int {
	var count int
	for i := range arr {
		for range arr[i] {
			count++
		}
	}
	return count
}
