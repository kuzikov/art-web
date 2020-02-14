package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CheckArgs func using for validation arguments
func CheckArgs(args []string, p *Params) {
	//func checkArgs(args []string, initColorState, initBgColorState, fileToOpen, fileToOutput, strInput, font string, output, reverse, strInputBool bool) {
	if len(args) == 0 {
		fmt.Println("Ascii-art consists on receiving a string has an argument\nand outputting the string in a graphic representation of ASCII.\nThis project should handle numbers, letters, spaces, special characters and \\n.")
		fmt.Println("You can use different flags with this program;\nalso you can change font of the text\ntype of fonts:\n   1) standard;\n   2) shadow;\n   3) thinkertoy;")
		fmt.Println("color of fonts --color=:\n   1) red;\n   2) blue;\n   3) green;\n   4) white;\n   5) orange;\n   6) yellow;")
		fmt.Println("color of background --bg=:\n   1) red;\n   2) blue;\n   3) green;\n   4) white;\n   5) orange;\n   6) yellow;")
		fmt.Println("align of the text --align=:\n   1) center;\n   2) right;\n   3) left;\n   4) justify;")
		fmt.Println("you can save output to file --output=<FILENAME.TXT>;")
		fmt.Println("you can open file and read them and print like NORMAL TEXT --reverse=<FILENAME.TXT>;")
		os.Exit(0)
	}

	///check flags///
	for i := range args {
		if !p.IsCorrectInput {
			// if strings.SplitN(args[i], "=", 2)[0] == "--reverse" {
			// 	goto here
			// }
			p.StrInput = args[i]
			// if !ValidInput(p.StrInput) {
			// 	fmt.Println("Incorrect input, change your language")
			// 	os.Exit(0)
			// }
			p.IsCorrectInput = true
			continue
		}
		// here:
		if args[i][0] == '-' {
			tempFlag := strings.SplitN(args[i], "=", 2)
			if len(tempFlag) == 1 {
				fmt.Println("flag is incorrect:\n you can use this flags:\n --color=<color> - to change color of your font;\n --align=<align> - to change offset of your printable text;\n --bg=<color> - to change background color;\n --output=<fileName.txt> - to create file with ASCII text;\n --reverse=<fileName.txt> - to open this file, read them and show in terminal all symbols from this file like NORMAL TEXT;")
				os.Exit(0)
			}

			switch tempFlag[0] {
			case "--align":
				p.Align = true
				switch tempFlag[1] {
				case "center":
					p.Textpos = "center"
				case "left":
					p.Textpos = "left"
				case "right":
					p.Textpos = "right"
				case "justify":
					p.Textpos = "justify"
				default:
					fmt.Println("flag --align is incorrect:\nyou must choose one of this offsets:\n 'center' --align=center\n 'left' --align=left\n 'right' --align=right\n 'justify' --align=justify\nor by default offset will be left")
					os.Exit(0)
				}
			case "--output":
				p.Output = true
				p.FileToOutput = tempFlag[1]
			case "--color":
				p.Color = true
				switch tempFlag[1] {
				case "red", "RED", "r", "R":
					p.InitColorState = FgRed
				case "green", "GREEN", "g", "G":
					p.InitColorState = FgGreen
				case "blue", "BLUE", "b", "B":
					p.InitColorState = FgBlue
				case "yellow", "YELLOW", "y", "Y":
					p.InitColorState = FgYellow
				case "orange", "ORANGE", "o", "O":
					p.InitColorState = FgOrange
				case "purple", "PURPLE", "p", "P":
					p.InitColorState = FgPurple
				case "cyan", "CYAN", "c", "C":
					p.InitColorState = FgCyan
				case "white", "WHITE", "w", "W":
					p.InitColorState = FgWhite
				default:
					fmt.Println("flag --color is incorrect:\nyou must choose one of this colors:\n 'red' --color=red\n 'yellow' --color=yellow\n 'white' --color=white\n 'green' --color=green\n 'blue' --color=blue")
					os.Exit(0)
				}
			case "--bg":
				p.Color = true
				switch tempFlag[1] {
				case "red", "RED", "r", "R":
					p.InitBgColorState = BgRed
				case "green", "GREEN", "g", "G":
					p.InitBgColorState = BgGreen
				case "blue", "BLUE", "b", "B":
					p.InitBgColorState = BgBlue
				case "yellow", "YELLOW", "y", "Y":
					p.InitBgColorState = BgYellow
				case "orange", "ORANGE", "o", "O":
					p.InitBgColorState = BgOrange
				case "purple", "PURPLE", "p", "P":
					p.InitBgColorState = BgPurple
				case "cyan", "CYAN", "c", "C":
					p.InitBgColorState = BgCyan
				case "white", "WHITE", "w", "W":
					p.InitBgColorState = BgWhite
				default:
					fmt.Println("flag --bg is incorrect:\nyou must choose one of this colors:\n 'red' --bg=red\n 'yellow' --bg=yellow\n 'white' --bg=white\n 'green' --bg=green\n 'blue' --bg=blue")
					os.Exit(0)
				}
			case "--reverse":
				p.Reverse = true
				if tempFlag[1] == "" {
					fmt.Println("error: no file to open")
					os.Exit(0)
				}
				p.FileToOpen = tempFlag[1]
				break
			case "--pipe":
				if tempFlag[1] == "true" || tempFlag[1] == "on" || tempFlag[1] == "ON" {
					p.Pipe = true
				}
			case "--colorSlice":
				//colorSlice=TEXT=COLOR
				colorSlice := strings.Split(tempFlag[1], "=")
				if colorSlice[0] == "" {
					fmt.Println("incorrect slice flag: use <<--colorSlice=TEXT=COLOR>>")
					os.Exit(0)
				}
				p.SliceBool = true
				p.Slice = colorSlice[0]
				switch colorSlice[1] {
				case "red", "RED", "r", "R":
					p.ColorOfSlice = FgRed
				case "green", "GREEN", "g", "G":
					p.ColorOfSlice = FgGreen
				case "blue", "BLUE", "b", "B":
					p.ColorOfSlice = FgBlue
				case "yellow", "YELLOW", "y", "Y":
					p.ColorOfSlice = FgYellow
				case "orange", "ORANGE", "o", "O":
					p.ColorOfSlice = FgOrange
				case "purple", "PURPLE", "p", "P":
					p.ColorOfSlice = FgPurple
				case "cyan", "CYAN", "c", "C":
					p.ColorOfSlice = FgCyan
				case "white", "WHITE", "w", "W":
					p.ColorOfSlice = FgWhite
				default:
					fmt.Printf("%s", "u can use this colors:\n RED =red, =RED, =r, =R\n GREEN =green, =GREEN, =G, =g\n BLUE =blue, =BLUE, =B, =b\n YELLOW =yellow, =YELLOW, =Y, =y\n ORANGE =orange, =ORANGE, =O, =o\n")
					os.Exit(0)
				}
			case "--colorInterval":
				//--colorInteval=int:int=COLOR
				tempArray := strings.SplitN(tempFlag[1], "=", 2)
				if len(tempArray) == 1 {
					fmt.Println("flag Colorized interval is incorrect. use : <<--colorInterval=INT:INT=color")
					os.Exit(0)
				}
				tempArray2 := strings.SplitN(tempArray[0], ":", 2)
				if len(tempArray2) == 1 {
					fmt.Println("flag Colorized interval is incorrect. use : <<--colorInterval=INT:INT=color")
					os.Exit(0)
				}
				start, err := strconv.Atoi(tempArray2[0])
				finish, err2 := strconv.Atoi(tempArray2[1])
				if err != nil || err2 != nil {
					fmt.Println("flag Colorized interval is incorrect. use : <<--colorInterval=INT:INT=color")
					os.Exit(0)
				}
				if start <= 0 || finish <= 0 || start > finish {
					fmt.Println("coordinates of flag is not corrected. INT1 > 0 & INT2>=INT1")
					os.Exit(0)
				}
				intArray := []int{}
				for i := start; i <= finish; i++ {
					intArray = append(intArray, i)
				}
				p.Interval = true
				p.IntArr = intArray
				switch tempArray[1] {
				case "red", "RED", "r", "R":
					p.ColorOfInterval = FgRed
				case "green", "GREEN", "g", "G":
					p.ColorOfInterval = FgGreen
				case "blue", "BLUE", "b", "B":
					p.ColorOfInterval = FgBlue
				case "yellow", "YELLOW", "y", "Y":
					p.ColorOfInterval = FgYellow
				case "orange", "ORANGE", "o", "O":
					p.ColorOfInterval = FgOrange
				case "purple", "PURPLE", "p", "P":
					p.ColorOfInterval = FgPurple
				case "cyan", "CYAN", "c", "C":
					p.ColorOfInterval = FgCyan
				case "white", "WHITE", "w", "W":
					p.ColorOfInterval = FgWhite
				default:
					fmt.Printf("%s", "u can use this colors:\n RED =red, =RED, =r, =R\n GREEN =green, =GREEN, =G, =g\n BLUE =blue, =BLUE, =B, =b\n YELLOW =yellow, =YELLOW, =Y, =y\n ORANGE =orange, =ORANGE, =O, =o\n")
					os.Exit(0)
				}
			case "--colorChar":
				//--colorChar=RUNE=COLOR
				tempArray := strings.SplitN(tempFlag[1], "=", 2)
				if len(tempArray) == 1 {
					fmt.Println("flag Colorized Char is incorrect. use : <<--colorInterval=CHAR=color")
					os.Exit(0)
				}
				if len(tempArray[0]) != 1 {
					fmt.Println("u can write only one symbol to color him")
					os.Exit(0)
				}
				p.Char = true
				char := &Char{Symbol: rune(tempArray[0][0])}
				switch tempArray[1] {
				case "red", "RED", "r", "R":
					char.Color = FgRed
				case "green", "GREEN", "g", "G":
					char.Color = FgGreen
				case "blue", "BLUE", "b", "B":
					char.Color = FgBlue
				case "yellow", "YELLOW", "y", "Y":
					char.Color = FgYellow
				case "orange", "ORANGE", "o", "O":
					char.Color = FgOrange
				case "purple", "PURPLE", "p", "P":
					char.Color = FgPurple
				case "cyan", "CYAN", "c", "C":
					char.Color = FgCyan
				case "white", "WHITE", "w", "W":
					char.Color = FgWhite
				default:
					fmt.Println("incorrect color of char")
					os.Exit(0)
				}
				p.CharArr = append(p.CharArr, *char)
			default:
				fmt.Println("flag is incorrect:\n you can use this flags:\n --color=<color> - to change color of your font;\n --align=<align> - to change align of your printable text;\n --bg=<color> - to change background color;\n --output=<fileName.txt> - to create file with ASCII text;\n --reverse=<fileName.txt> - to open this file, read them and show in terminal all symbols from this file like NORMAL TEXT;")
				os.Exit(0)
			}
		} else {
			p.Font = args[i] + ".txt"
		}
	}
}

// ValidInput scan symbols in arguments for incorrectable symbols
func ValidInput(arr string) bool {
	for i := range arr {
		if arr[i] == 13 {
			continue
		}
		if arr[i] < 32 || arr[i] > 126 {
			return false
		}
	}
	return true
}
