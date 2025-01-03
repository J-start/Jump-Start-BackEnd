package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

/*

	Source code from: https://gist.github.com/fracasula/38aa1a4e7481f9cedfa78a0cdd5f1865
	
	Credits To: Francesco Casula

*/

func EncryptMessage(key []byte, message string) (string, error) {
	//Key must be 16,24 or 32 bytes long (AES-128,192,256)

	byteMsg := []byte(message)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.New("erro ao criar novo cipher")
	}

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", errors.New("erro ao criar novo cipher")
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptMessage(key []byte, message string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", errors.New("erro ao descrifrar mensagem")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.New("erro ao criar NewCipher")
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("erro ao criar NewCipher")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

