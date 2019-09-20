package main

import (
	"fmt"

	cp "./cipher"
)

func main() {

	var cipher cp.ICipher
	cipher = cp.ReplacementCipher{}
	key := cipher.GenerateKey(8)
	encoded := cipher.Encode(key, "aaaa")
	decoded := cipher.Decode(key, encoded)
	fmt.Println("key: ", key, " text: ", encoded, " Decode:", decoded)
}
