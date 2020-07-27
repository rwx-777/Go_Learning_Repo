package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	fmt.Println("Creating Hash...")
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	createHash("Carrie")
}
