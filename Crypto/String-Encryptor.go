package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func createHashInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter String to Hash: ")
	key, _ := reader.ReadString('\n')
	hasher := md5.New()
	hasher.Write([]byte(key))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func encryptString() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Text to Encrypt: ")
	textToEncrypt, _ := reader.ReadString('\n')

	ciphertext := encrypt([]byte(textToEncrypt), "1234")
	fmt.Println(ciphertext)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func decryptString() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter cipher bytes to decrypt: ")
	ciphertext, _ := reader.ReadBytes('\n')
	fmt.Println("Enter password used for encryption: ")
	passphrase, _ := reader.ReadString('\n')

	plaintext := decrypt(ciphertext, passphrase)
	fmt.Printf("Decrypted: %s\n", plaintext)
	return string(plaintext)
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func encryptFileName(filename string, passphrase string) {

	//os.Rename(encrypt(filename, data))
}

func printMenu() {

	fmt.Println("1. Hash String MD5")
	fmt.Println("2. Encrypt String")
	fmt.Println("3. Decrypt String")
	fmt.Println("4. Encrypt Filename")
	fmt.Println("5. Decrypt Filename")
	fmt.Println("6. Encrypt File Contents")
	fmt.Println("7. Decrypt File Contents")

}

func main() {

	fmt.Println("Starting Application...")
	printMenu()
	fmt.Println("Select your Option: ")
	var i int
	fmt.Scan(&i)

	switch {

	case i == 1:
		createHashInput()
	case i == 2:
		encryptString()
	case i == 3:
		decryptString()
	case i == 6:

	case i == 7:

	default:
		fmt.Println("No Input")
	}

	//createHashInput()

	//fmt.Println("Starting the application...")
	//ciphertext := encrypt([]byte("Hello World"), "password")
	//fmt.Printf("Encrypted: %x\n", ciphertext)
	//plaintext := decrypt(ciphertext, "password")
	//fmt.Printf("Decrypted: %s\n", plaintext)
	//encryptFile("sample.txt", []byte("Hello World"), "password1")
	//fmt.Println(string(decryptFile("sample.txt", "password1")))
}
