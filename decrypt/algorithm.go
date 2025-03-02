package decrypt

func Nimbus(str string) string {
	decryptedStr := ""

	for _, c := range str {
		asciicode := int(c)
		character := string(asciicode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
