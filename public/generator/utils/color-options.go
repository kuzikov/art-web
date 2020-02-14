package utils

import "fmt"

//ColorSlice take color params for specific slice
func ColorSlice(str, slice, color string, index int, indexofSlice *int, tempBool *bool) {
	if !*tempBool {
		if str[index] == slice[*indexofSlice] {
			if index <= len(str)-len(slice) {
				tempStr := str[index : index+len(slice)]
				if tempStr == slice {
					*tempBool = true
					*indexofSlice++
					/////
					fmt.Printf("%v", SRG(color))
					/////
				}
				if *indexofSlice == len(slice) {
					*indexofSlice = 0
					*tempBool = false
				}
			}
		}
	} else {
		if str[index] == slice[*indexofSlice] {
			//////
			fmt.Printf("%v", SRG(color))
			//////
			*indexofSlice++
			if *indexofSlice == len(slice) {
				*indexofSlice = 0
				*tempBool = false
			}
		}
	}
}

//ColorInterval take color for specified interval
//of positions
func ColorInterval(index int, arr []int, color string) {
	for i := range arr {
		if arr[i]-1 == index {
			fmt.Printf("%v", SRG(color))
		}
	}
}

//ColorChar take color for specified Symbol
func ColorChar(r rune, arr []Char) {
	for i := range arr {
		if arr[i].Symbol == r {
			fmt.Printf(SRG("6"))
			fmt.Printf("%v", SRG(arr[i].Color))
		}
	}
}
