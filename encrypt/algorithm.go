package encrypt

func Nimbus(str string) string {
	encryptedStr := ""

	for _, c := range str {
		asciicode := int(c)
		character := string(asciicode + 3)
		encryptedStr += character
	}
	return encryptedStr
}
