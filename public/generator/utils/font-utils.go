package utils

import (
	"os"
)

// ChooseFont func open font file like array of string
func ChooseFont(p *Params) (arr []string, hight int, message string) {
	////check font from file & letter hight
	file, err := os.Open("./public/generator/files/" + p.Font)
	if err != nil {
		return []string{}, 0, "Error: file<" + p.Font + "> not found"
	}
	f, _ := file.Stat()
	size := int64(f.Size())
	array := make([]byte, size)
	file.Read(array)
	////create array of string
	arrString := []string{}
	tempStr := ""
	//////ascii symbol height
	lines := 0
	separatorHight := 0
	height := false
	////////////
	for i := range array {
		if array[i] == 10 && tempStr == "" && height == false {
			separatorHight++
		}
		if array[i] == 10 && tempStr != "" && height == false {
			height = true
		}
		if array[i] == 10 {
			if tempStr != "" {
				arrString = append(arrString, tempStr)
				tempStr = ""
			}
			if separatorHight > 0 && tempStr != "" {
				height = true
			}
			lines++
			continue
		}
		tempStr = tempStr + string(array[i])
	}
	if lines%95 != 0 {
		return []string{}, 0, "Error: file is incorrect"
	}
	letterHight := lines/95 - separatorHight
	if tempStr != "" {
		arrString = append(arrString, tempStr)
		tempStr = ""
	}
	return arrString, letterHight, ""
}
