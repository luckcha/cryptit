package main

import (
	"fmt"

	"github.com/luckcha/decrypt"
	"github.com/luckcha/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Infracloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))

}
