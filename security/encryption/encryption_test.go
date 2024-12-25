package encryption

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const isBase64 = "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$"
/*

	Source code from : https://gist.github.com/fracasula/38aa1a4e7481f9cedfa78a0cdd5f1865
	
	Credits To : Francesco Casula

*/
func TestEncryptDecryptMessage(t *testing.T) {
	key := []byte("0123456789abcdef") 
	message := "Lorem ipsum dolor sit amet"

	encrypted, err := EncryptMessage(key, message)
	fmt.Println(encrypted)
	require.Nil(t, err)
	require.Regexp(t, isBase64, encrypted)

	decrypted, err := DecryptMessage(key, encrypted)
	fmt.Println(decrypted)	
	require.Nil(t, err)
	require.Equal(t, message, decrypted)
}