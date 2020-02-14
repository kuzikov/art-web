package utils

import (
	"fmt"
	"os"
)

// ReadFromFile func using when flag <--reverse=> activated
func ReadFromFile(p *Params) {
	//opening needed file
	f, e := os.Open(p.FileToOpen)
	if e != nil {
		fmt.Printf("%s", e.Error())
		fmt.Println()
		return
	}
	f2, _ := f.Stat()
	size := int64(f2.Size())
	array := make([]byte, size)
	f.Read(array)
	arrStr := []string{}
	tempStr := ""
	//lH -- Hight of letter (by default is 8)
	lH := 8
	lines := 0
	//checking for using "\n" and splitting input file by different array
	for i := range array {
		if array[i] == '\n' {
			lines++
			if lines%lH == 0 {
				arrStr = append(arrStr, tempStr)
				tempStr = ""
			}
			continue
		}
		tempStr = tempStr + string(rune(array[i]))
	}
	//if tempory string is not empty at finalized operation - that file is incorrect
	if tempStr != "" {
		fmt.Println("Input file is INCORRECT")
		return
	}
	// make stack for pull them from strings from new ^^^ array and compare him with tempory stack constucted from symbols
	stack := make([]string, 8)
	inputString := ""
	finalOutput := []string{}
	for i := range arrStr {
		lenStr := len(arrStr[i]) / lH
		for k := 0; k < lenStr; k++ {
			for j := 0; j < lH; j++ {
				stack[j] = stack[j] + string(rune(arrStr[i][k+j*lenStr]))
			}
			tempRune, err := CheckStack(stack, p)
			if err == true {
				///////add func with special symbols!!!!!!!
				// if tempRune == '`' {
				// 	inputString = inputString + "\\"
				// }
				inputString = inputString + string(tempRune)
				stack = make([]string, 8)
			}
		}
		//if inputFile will be incorrect - stack will be not empty
		if CompareStack(stack, make([]string, 8)) {
			if i != len(arrStr)-1 {
				//inputString = inputString + "\\n"
				finalOutput = append(finalOutput, inputString)
				inputString = ""
			}
		} else {
			fmt.Println("Error: Input file has incorrect strings")
			return
		}

	}
	finalOutput = append(finalOutput, inputString)
	//print to termilal finalized string
	//fmt.Println(finalOutput)
	for i := range finalOutput {
		fmt.Printf("%s\n", finalOutput[i])
	}
}
