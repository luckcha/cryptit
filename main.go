package main

import (
	"fmt"

	"github.com/luckcha/cryptit/decrypt"
	"github.com/luckcha/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Infracloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))

}
