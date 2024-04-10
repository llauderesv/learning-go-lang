package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "He is Risen, Happy Easter Sunday!"

	hash256 := sha256.Sum256([]byte(str))

	hashString := hex.EncodeToString(hash256[:])

	fmt.Println("SHA-256 hash: ", hashString)
	fmt.Println("SHA-256 decode hash: ", hashString)

	hash384 := sha512.Sum384([]byte(str))
	hash384String := hex.EncodeToString(hash384[:])
	fmt.Println("SHA-384 hash: ", hash384String)

	hash512 := sha512.Sum512([]byte(str))
	hash512String := hex.EncodeToString(hash512[:])
	fmt.Println("SHA-512 hash: ", hash512String)
}
