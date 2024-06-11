package Ascii

func CheckAlphabets(text string) bool {
	for _, char := range text {
		if char < 32 && char != 10 && char != 13 || char > 126 {
			return false
		}
	}
	return true
}
