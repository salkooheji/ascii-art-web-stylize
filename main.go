package main

import (
	Ascii "Ascii/functions"
	"fmt"
)

func main() {
	// //extracting the file name from the args
	// var shapeFileContent string
	// // check what's the defualt ?
	// // if len(os.Args)==4 {
	// // 	if len(os.Args) == 2 { // standard is the default
	// // 		shapeFileContent = functions.ReadShapeFile("standard.txt")
	// // 	} else

	// if len(os.Args) == 4 {
	// 	if functions.ValidateArugments() {
	// 		shapeFileContent = functions.ReadShapeFile(os.Args[2])
	// 	} else {
	// 		log.Fatal("Check the input entered the program expect 3 values (Value to be displayed text (string), Font (.txt), the file to show the result in (.txt))")
	// 	}
	// } else {
	// 	log.Fatal("Invalid Number of arguments")
	// }

	// charactersMap := functions.MapFileContentToRunes(os.Args[1], shapeFileContent)
	// textToDraw := os.Args[1]
	// // This will help in handling \n each index in textToDrawSplit will be considered a line
	// textToDrawSplit := strings.Split(textToDraw, `\n`)
	// result := ""
	// for i, line := range textToDrawSplit {
	// 	if line != "" {
	// 		generateShape := functions.GenerateShape(line, charactersMap)
	// 		//fmt.Println(generateShape)
	// 		result += generateShape

	// 	}

	// 	if i < len(textToDrawSplit)-1 {
	// 		result += "\n"
	// 	}
	// }
	// os.WriteFile(os.Args[3], []byte(result), 0o644)
	fmt.Println(Ascii.FinalShape())
}
