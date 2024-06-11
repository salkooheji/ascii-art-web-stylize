package Ascii

import "strings"

func FinalResult(text string, font string) string {
	font = "fonts/" + font
	shapeFileContent := ReadShapeFile(font)
	charactersMap := MapFileContentToRunes(shapeFileContent)
	textToDrawSplit := strings.Split(text, "\n")
	result := ""

	for _, line := range textToDrawSplit {
		if line != "" {
			generateShape := GenerateShape(line, charactersMap)
			result += generateShape + "\n"
		}
	}

	return result
}
