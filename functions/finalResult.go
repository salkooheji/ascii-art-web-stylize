package Ascii

import "strings"

func FinalResult(text string, font string) string {
	font = "fonts/" + font
	shapeFileContent := ReadShapeFile(font)
	charactersMap := MapFileContentToRunes(shapeFileContent)
	
	if text == "" {
		return ""
	}
	textToDrawSplit := strings.Split(text, "\n")
	result := ""

	if IsEmpty(textToDrawSplit) {
		textToDrawSplit = textToDrawSplit[1:]
	}

	for _, line := range textToDrawSplit {
		if line != "" {
			generateShape := GenerateShape(line, charactersMap)
			result += generateShape + "\n"
		}
	}

	return result
}

func IsEmpty(arr []string) bool {
	for _, v := range arr {
		if v != "" {
			return false
		}
	}
	return true
}
